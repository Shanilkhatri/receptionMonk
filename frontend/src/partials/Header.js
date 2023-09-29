import React, { useState, useEffect } from "react";
import logo from "../assets/images/logo_rm.svg";

function Header() {

    const [scrolled, setScrolled] = useState(false);

    useEffect(() => {
        const handleScroll = () => {
            if (window.scrollY > 50) {
                setScrolled(true);
            } else {
                setScrolled(false);
            }
        };

        window.addEventListener('scroll', handleScroll);

        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, {});

    const [collapsed, setCollapsed] = useState(true);

    const toggleNavbar = () => {
        setCollapsed(!collapsed);
      };

    return (
        <div>
            <nav
                style={{ backgroundColor: scrolled ? 'white' : 'transparent',
                boxShadow: scrolled ? '0 2px 4px rgba(0, 0, 0, 0.1)' : 'none' }}
                className={`navbar navbar-expand-lg fixed-top ${scrolled ? 'scrolled' : ''}`}>

                <div className="container">                                         
                    <a className="navbar-brand" href="#">
                        <img src={logo} alt="logo"/>
                    </a>                        
                    <button className="navbar-toggler" type="button" onClick={toggleNavbar}>
                        <span className="navbar-toggler-icon"></span>
                    </button>     
                
                    <div className={`collapse navbar-collapse ${collapsed ? '' : 'show'}`} id="navbarCollapse">
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
                            <li className="nav-button p-2 ms-lg-0 ms-xl-5">
                                <button className="btn btn-primary">Portal</button>
                            </li>
                        </ul>
                    </div>                        
                </div>
            </nav>
        </div>
    );
}

export default Header;