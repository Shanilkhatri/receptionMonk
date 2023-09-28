import React from "react";
import Header from "../partials/Header";
import Hero from "../components/Hero";
import ContentOne from "../components/BlockContentOne";
import ContentTwo from "../components/BlockContentTwo";
import BlockPrice from "../components/PriceTable";
import Footer from "../partials/Footer";

function homePage() {
    return (
        <div>
            
            <div className="bg-xs">
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
            </div>
               


            <div>

                {/* Price Section */}
                <BlockPrice />

            </div>   

            <div>

                {/* Footer Section */}
                <Footer />
            </div>

        </div>    
    );
}
export default homePage;