class Utility{
    appUrl: string;

    constructor() {
        this.appUrl = import.meta.env.VITE_APPURL;

    }


getCookieValue(key:any) {
        const cookies = document.cookie.split(';');
        
        for (let i = 0; i < cookies.length; i++) {
            const cookie = cookies[i].trim();
            
            // Check if the cookie starts with the provided key
            if (cookie.startsWith(key + '=')) {
                // Extract and return the cookie value
                return cookie.substring(key.length + 1);
            }
        }
        
        // If the cookie with the given key is not found
        return null;
    }
async sendRequestPutOrPost(userData:any,route:string,method:string) {

    const response = await fetch(this.appUrl+route, {
        method: method,
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${this.getCookieValue("exampleToken")}`
        },
        body: JSON.stringify(userData),
    });
    // var responseData = await response.json() // wait for response > json
    if (!response.ok) {
        // navigate("/auth/SignIn")
        return false
    } else if (response.ok) {
        return true
    }
    return false
    }
    async uploadImageAndReturnUrl(
        imageElement: string,
        modulename: string,
        route: string
      ): Promise<string> {
        const imageUploadURL = this.appUrl + route;
    
        const imageInput = document.getElementById(
          imageElement
        ) as HTMLInputElement | null;
        let imageFile: File | undefined;
    
        if (
          imageInput !== null &&
          imageInput.files &&
          imageInput.files.length > 0
        ) {
          imageFile = imageInput.files[0];
        }
        const formData = new FormData();
        if (imageFile != undefined && imageInput != null) {
          const valid = this.validateImgBeforeUpload(imageFile);
          if (!valid) {
            console.error("not a valid image is uploaded:");
            return "";
          }
          formData.append("image", imageFile || "");
          formData.append("modulename", modulename);
        }
        try {
          var response = await fetch(imageUploadURL, {
            method: "POST",
            headers: {
              Authorization: `Bearer ${this.getCookieValue("exampleToken")}`,
            },
            body: formData,
          });
    
          if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
          }
          const data = await response.json();
          return data.Payload;
        } catch (error) {
          console.error("Error during image upload:", error);
          return ""; // or throw an error, depending on your logic
        }
      }
     validateImgBeforeUpload(imageFile: File): boolean {
        var allowedExtensions = ["jpg", "jpeg", "png"]; // Allowed image extensions
        var maxFileSize = 5 * 1024 * 1024; // Maximum file size in bytes (5MB)
        if (!imageFile.name) {
          console.error("File name is undefined");
          return false;
        }
        // Check the file extension
        const fileParts = imageFile.name.split(".");
        if (fileParts.length === 1) {
          console.error("File name does not have an extension");
          return false;
        }
        // Check the file extension
        var fileExtension = fileParts.pop()!.toLowerCase();
        if (!allowedExtensions.includes(fileExtension)) {
          return false;
        }
        // Check the file size
        if (imageFile.size > maxFileSize) {
          console.log("failure", "Image size should be under 5mb");
          return false;
        }
        return true;
      }
}

export default Utility