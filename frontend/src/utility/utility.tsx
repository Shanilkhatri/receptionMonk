
class Utility{

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

    const response = await fetch("http://localhost:4000/"+route, {
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
}

export default Utility