import React from "react";
import Header from "../components/Header";
import Hero from "../components/Hero";
import ContentOne from "../components/BlockContentOne";

function homePage() {
    return (
        <div>
            
            <div className="bg-anim-one"> 

                {/* Header Section */}
                <Header />     

                {/* Hero Section */}
                <Hero /> 

            </div>
            
           <div className="bg-anim-two"> 

                {/* Content Block One Section */}
                <ContentOne />

            </div>         

        </div>    
    );
}
export default homePage;