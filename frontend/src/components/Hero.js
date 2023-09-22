import React from "react";
import anim from "../assets/images/scrnshot.png";
import sponcer from "../assets/images/sponcer.svg";

function Hero() {
    return (
        <div>
            <section id="hero">
                <div className="container">
                    <div className="row align-items-center hero-content">
                        <div className="col-md-6">
                            <div className="hero-text text-primary">
                                <h1 className="fw-bold mb-4 display-3">Take your Team to the next level</h1>
                                <p className="my-3">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                            </div>
                            <div>
                                <button className="btn btn-primary mt-5">CTA</button>
                            </div>
                        </div>
                        <div className="col-md-6">
                            <div className="anim">
                                <img src={anim} alt="" className="img-fluid" />
                            </div>
                        </div>
                    </div>
                    <div className="sponcership">
                        <div className="row">
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
                </div>
            </section>
        </div>
    );
}

export default Hero;