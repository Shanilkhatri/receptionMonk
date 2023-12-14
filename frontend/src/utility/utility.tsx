import Swal from "sweetalert2";

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
  // One function for all the requests on "user route" 
  // handles PUT POST GET
  // Send data in for a JSON Object {} for PUT/POST
  // Send an empty string in case of GET
  // It returns the whole response from the server
  // response structure has keys : 
  // response = {
  //   // "Status": // indicating status for the request
  //   // "Message": // indicating message from the server against the request
  //   // "Payload": [{}] // contains the main data against the query
  // }
async sendRequest_Put_Post_Get(userData:any,route:string,method:string) {

  const requestOptions: RequestInit & { body?: string } = {
    method: method,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${this.getCookieValue("exampleToken")}`
    },
  };
  // Check if the method is one that allows a request body
  if (['POST', 'PUT', 'PATCH'].includes(method.toUpperCase())) {
    // Assuming userData is defined
    requestOptions.body = JSON.stringify(userData);
  }
    const response = await fetch(this.appUrl+route, requestOptions);
    
    var responseData = await response.json() // wait for response > json
   if (response.ok) {
      const toast = Swal.mixin({
        toast: true,
        position: "top-end",
        showConfirmButton: false,
        timer: 3000,
      });
      toast.fire({
        icon: "success",
        title: responseData.Message,
        padding: "10px 20px",
      });
      return responseData
    }
    const toast = Swal.mixin({
      toast: true,
      position: "top-end",
      showConfirmButton: false,
      timer: 3000,
    });
    toast.fire({
      icon: "error",
      title: responseData.Message,
      padding: "10px 20px",
    });
    return responseData
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