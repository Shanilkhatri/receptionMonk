import React from "react";
import Header from "../components/Header";
import Hero from "../components/Hero";
import ContentOne from "../components/BlockContentOne";
import ContentTwo from "../components/BlockContentTwo";
import BlockPrice from "../components/PriceTable";

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

                {/* Content Block Two Section */}
                <ContentTwo />

            </div>      

            <div>

                {/* Price Section */}
                <BlockPrice />

            </div>   

        </div>    
    );
}
export default homePage;