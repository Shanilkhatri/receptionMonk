import React from "react";
import anim from "../assets/images/scrnshot.jpeg";
import sponcer from "../assets/images/sponcer.svg";

function Hero() {
    return (
        <div>
            <section id="hero">
                <div className="container">
                    <div className="row align-items-center">
                        <div className="col-md-6">
                            <div className="hero-text text-primary">
                                <h1 className="fw-bold mb-3 mb-lg-4 display-3">Take your Team to the next level</h1>
                                <p className="my-1 my-lg-3">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                            </div>
                            <div>
                                <button className="btn btn-primary mt-3 mt-lg-5 mb-4 mb-md-0">CTA</button>
                            </div>
                        </div>
                        <div className="col-md-6">
                            <div className="anim">
                                <img src={anim} alt="" className="img-fluid" />
                            </div>
                        </div>
                    </div>
                    <div className="row align-items-center d-none d-md-flex sponcership">                        
                        <div className="col-md-2">
                            <img src={sponcer} alt="clients logo" className="img-fluid"></img>
                        </div>
                        <div className="col-md-2">
                            <img src={sponcer} alt="clients logo" className="img-fluid"></img>
                        </div>
                        <div className="col-md-2">
                            <img src={sponcer} alt="clients logo" className="img-fluid"></img>
                        </div>
                        <div className="col-md-2">
                            <img src={sponcer} alt="clients logo" className="img-fluid"></img>
                        </div>
                        <div className="col-md-2">
                            <img src={sponcer} alt="clients logo" className="img-fluid"></img>
                        </div>
                    </div>                    
                </div>
            </section>
        </div>
    );
}

export default Hero;