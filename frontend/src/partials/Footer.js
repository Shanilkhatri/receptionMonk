import React from "react";

function Footer() {
    return (
        <div>
            <footer>
                <section id="footer" className="bg-primary">
                    <div className="container">
                        <div className="row">
                            <div className="col-md-4 d-flex justify-content-center">
                                <div>
                                    <h3 className="text-white d-none d-md-block">Reception <br/> Monk</h3>
                                    <h3 className="text-white d-block d-md-none">Reception Monk</h3>
                                    <div className="d-flex social mt-3 justify-content-center justify-content-md-start">
                                        <div><i className="bi bi-facebook"></i></div>
                                        <div className="mx-2"><i className="bi bi-instagram border-end border-start border-white px-2"></i></div>
                                    </div>
                                </div>                               
                            </div>
                            <div className="col-md-4">
                                <div className="d-flex justify-content-start justify-content-md-center">
                                    <ul className="list-group">
                                        <li className="list-group-item">About us</li>
                                        <li className="list-group-item">Contact us</li>
                                        <li className="list-group-item">Promos</li>
                                        <li className="list-group-item">Portal</li>                                   
                                    </ul>
                                </div>
                            </div>
                            <div className="col-md-4">
                                <div className="d-flex justify-content-start justify-content-md-center">
                                    <ul className="list-group">
                                        <li className="list-group-item">Terms of Usage</li>
                                        <li className="list-group-item">Privacy Polidy</li>
                                    </ul>
                                </div>                               
                            </div>
                        </div>
                        <div className="row">
                            <div className="d-flex justify-content-center my-4">
                                <h6>Copyright &copy; 2023 | ReceptionMonk</h6>
                            </div>
                        </div>
                    </div>
                </section>
            </footer>
        </div>
    );
}

export default Footer;