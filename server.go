package main

import (
	"context"
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"reakgo/models"
	"reakgo/router"
	"reakgo/utility"

	"github.com/allegro/bigcache/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	// Set log configuration
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println(".env file wasn't found, looking at env variables")
	}
	var Helper utility.Helper = &utility.Utility{}
	motd()
	// Read Config
	utility.Db, err = sqlx.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Println("Wowza !, We didn't find the DB or you forgot to setup the env variables")
		panic(err)
	}
	utility.Store = sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_KEY")))
	utility.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 1,
		HttpOnly: true,
	}
	utility.View = cacheTemplates()
	// See "Important settings" section.
	utility.Db.SetConnMaxLifetime(time.Minute * 3)
	utility.Db.SetMaxOpenConns(10)
	utility.Db.SetMaxIdleConns(10)

	gob.Register(utility.Flash{})
	utility.LogFile = Helper.OpenLogFile() // open Logfile
}

func main() {
	// Initialize Caching
	cacheInit()
	// Generate cache as a go routine as to not halt operation,
	// Cache fail-safe is already implemented so will fetch from DB incase the cache is not populated
	go models.GenerateCache()
	csrf_secret_key := os.Getenv("CSRF_SECRET_KEY")
	if csrf_secret_key == "" {
		log.Fatal("Missing Env value CSRF_SECRET_KEY")
	}

	utility.CSRF = csrf.Protect([]byte(csrf_secret_key))

	mux := mux.NewRouter()
	// Serve static assets
	staticHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	mux.PathPrefix("/assets/").Handler(staticHandler)

	// Set up a file server to serve the uploads folder
	uploadsURL := "/uploads/" // URL path to access the uploads folder
	mux.PathPrefix(uploadsURL).Handler(http.HandlerFunc(staticHandlerUpload))
	mux.PathPrefix("/").HandlerFunc(handler)

	if os.Getenv("APP_IS") == "monolith" {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("WEB_PORT"), utility.CSRF(mux)))
	} else if os.Getenv("APP_IS") == "microservice" {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("WEB_PORT"), mux))
	}
}

func cacheTemplates() *template.Template {

	funcMap := template.FuncMap{
		// Only to be used for SAFE attributes, SAFE = Computer Generated and not USER DRIVEN
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		// Only to be used for SAFE HTML, SAFE = Computer Generated and not USER DRIVEN
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		// Only to be used for SAFE URLs, SAFE = Computer Generated and not USER DRIVEN
		"safeURL": func(s string) template.URL {
			return template.URL(s)
		},
	}

	templ := template.New("")
	templ.Funcs(funcMap)
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}

// restrict the upload function to get access by the specific path only
func staticHandlerUpload(w http.ResponseWriter, r *http.Request) {
	// Get the requested file name
	fileName := filepath.Base(r.URL.Path)

	// Construct the actual file path within the /assets/ directory
	filePath := filepath.Join("./uploads/", fileName)

	// Serve the requested file
	http.ServeFile(w, r, filePath)
}

func handler(w http.ResponseWriter, r *http.Request) {
	router.Routes(w, r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, emailVerfToken,Authorization,tokenPayload")
	w.Header().Set("Content-Security-Policy", "default-src 'self'")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE")
}

func motd() {
	logo := `
______ _____  ___   _   __
| ___ \  ___|/ _ \ | | / /
| |_/ / |__ / /_\ \| |/ / 
|    /|  __||  _  ||    \ 
| |\ \| |___| | | || |\  \
\_| \_\____/\_| |_/\_| \_/
                          
----------------------------
Application should now be accessible on port ` + os.Getenv("WEB_PORT") + `

`
	log.Println(logo)
}

func cacheInit() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 512,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}

	var err error

	utility.Cache, err = bigcache.New(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
}
