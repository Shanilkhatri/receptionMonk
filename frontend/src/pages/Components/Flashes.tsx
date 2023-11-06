import { Link } from 'react-router-dom';
import React from 'react';
import Swal from 'sweetalert2';
import withReactContent from 'sweetalert2-react-content';
import { useSelector } from 'react-redux';
import { IRootState } from '../../store';

const Flashes = () => {
    const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;

    const MySwal = withReactContent(Swal);
   
    const coloredToast = (color: string) => {
        const toast = Swal.mixin({
            toast: true,
            position: isRtl ? 'top-start' : 'top-end',
            showConfirmButton: false,
            timer: 3000,
            showCloseButton: true,
            customClass: {
                popup: `color-${color}`,
            },
        });
        toast.fire({
            title: 'Type notification here, according to your requirement',
        });
    };

    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Components
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Flashes</span>
                </li>
            </ul>
            <div className="pt-5 space-y-8">                
                <div className="grid lg:grid-cols-2 grid-cols-1 gap-6">     
                    
                    {/* Flashes with background color */}
                    <div className="panel lg:col-span-2">
                        <div className="flex items-center justify-between mb-5">
                            <h5 className="font-semibold text-lg dark:text-white-light">Flash Notification</h5>                           
                        </div>
                        <div className="mb-5">
                            <div className="grid grid-cols-2 sm:grid-cols-1 sm:flex items-center justify-center gap-2 colored-toast">
                                <div>
                                    <button type="button" className="btn btn-primary" onClick={() => coloredToast('primary')}>
                                        General Notifiation
                                    </button>
                                    <div id="primary-toast"></div>
                                </div>
                                <div>
                                    <button type="button" className="btn btn-secondary" onClick={() => coloredToast('secondary')}>
                                       Theme Notification
                                    </button>
                                    <div id="secondary-toast"></div>
                                </div>
                                <div>
                                    <button type="button" className="btn btn-success" onClick={() => coloredToast('success')}>
                                        Success Notifiation
                                    </button> 
                                    <div id="success-toast"></div>
                                </div>
                                <div>
                                    <button type="button" className="btn btn-danger" onClick={() => coloredToast('danger')}>
                                        Failure Notifiation
                                    </button>
                                    <div id="danger-toast"></div>
                                </div>
                                <div>
                                    <button type="button" className="btn btn-warning" onClick={() => coloredToast('warning')}>
                                        Alert Notifiation
                                    </button>
                                    <div id="warning-toast"></div>
                                </div>
                                <div>
                                    <button type="button" className="btn btn-info" onClick={() => coloredToast('info')}>
                                        Informative Notifiation
                                    </button>
                                    <div id="info-toast"></div>
                                </div>                                
                            </div>
                        </div>                       
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Flashes;