import React from "react";
import serviceOne from "../assets/images/service-one.svg";
import serviceTwo from "../assets/images/service-two.svg";
import serviceThree from "../assets/images/service-three.svg";
import serviceFour from "../assets/images/service-four.svg";
import arrowRight from "../assets/images/arrow-right-bold.svg";


function  BlockContentOne() {
    return (
        <div>
            <section id="contentfirst">
                <div className="container">
                    <div className="row">
                        <div className="col-md-8 col-xl-6">
                            <div className="content-top-title">                                
                                <h2 className="fw-bold">The best calling solution for your business</h2>
                                <p className="my-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
                                <div className="link"><a href="#" className="text-black my-4 fw-bolder">Learn about our success <img src={arrowRight} alt="clickhere" className="ms-4" /></a> </div>
                            </div>
                        </div>
                    </div>
                    <div className="services text-center">
                        <div className="row">
                            <div className="col-md-3">   
                                <div className="serviceanim">
                                    <img src={serviceOne} alt="services" className="img-fluid" />
                                </div>
                                <div className="content-bottom-title">
                                    <h6>Single Brand Identity</h6>
                                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                                    <a href="#" className="my-4 fw-bolder">More <i className="bi bi-arrow-right"></i> </a>
                                </div>
                            </div>
                            <div className="col-md-3">
                                <div className="serviceanim">
                                    <img src={serviceTwo} alt="services" className="img-fluid" />
                                </div>
                                <div className="content-bottom-title">
                                    <h6>Works from anywhere</h6>
                                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                                    <a href="#" className="my-4 fw-bolder">More <i className="bi bi-arrow-right"></i> </a>
                                </div>
                            </div>
                            <div className="col-md-3">
                                <div className="serviceanim">
                                    <img src={serviceThree} alt="services" className="img-fluid" />
                                </div>
                                <div className="content-bottom-title">
                                    <h6>One Number for All</h6>
                                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                                    <a href="#" className="my-4 fw-bolder">More <i className="bi bi-arrow-right"></i> </a>
                                </div>
                            </div>
                            <div className="col-md-3">
                                <div className="serviceanim">
                                    <img src={serviceFour} alt="services" className="img-fluid" />
                                </div>
                                <div className="content-bottom-title">
                                    <h6>Advanced AI IVR</h6>
                                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                                    <a href="#" className="my-4 fw-bolder">More <i className="bi bi-arrow-right"></i> </a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    );
}

export default BlockContentOne;