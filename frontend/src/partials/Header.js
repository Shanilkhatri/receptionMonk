import React from "react";
import logo from "../assets/images/logo_rm.svg";

function Header() {
    return (
        <div>
            <div id="navbar">
                <nav className="navbar navbar-expand-md navbar-white">
                    <div className="container">                       
                        <a className="navbar-brand" href="#">
                            <img src={logo} alt="logo"/>
                        </a>
                        <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
                           
                        </button>
                        <div className="collapse navbar-collapse" id="navbarCollapse">
                            <ul className="navbar-nav navbar-center mx-auto">
                                <li className="nav-item">
                                    <a data-scroll href="#" className="nav-link">About</a>
                                </li>
                                <li className="nav-item">
                                    <a data-scroll href="#" className="nav-link">Technology</a>
                                </li>
                                <li className="nav-item">
                                    <a data-scroll href="#" className="nav-link">Pricing</a>
                                </li>
                                <li className="nav-item">
                                    <a data-scroll href="#" className="nav-link">Contact</a>
                                </li>
                                <li className="nav-button p-2 ms-5">
                                    <button className="btn btn-primary">Portal</button>
                                </li>
                            </ul>

                            {/* <div className="nav-button">
                                <button className="btn btn-primary">Portal</button>
                            </div> */}
                        </div>
                    </div>
                </nav>
            </div> 
        </div>
    );
}

export default Header;