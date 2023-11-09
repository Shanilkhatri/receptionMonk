
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
}

export default Utility