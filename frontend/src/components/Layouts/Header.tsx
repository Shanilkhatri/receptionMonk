import { useEffect, Fragment, useState } from 'react';
import { Dialog, Transition } from '@headlessui/react';
import { useDispatch, useSelector } from 'react-redux';
import { Link, NavLink, useLocation, useNavigate } from 'react-router-dom';
import { IRootState } from '../../store';
import { toggleTheme } from '../../store/themeConfigSlice';
import { useTranslation } from 'react-i18next';
import { toggleSidebar } from '../../store/themeConfigSlice';
import Dropdown from '../Dropdown';

const Header = () => {
    const location = useLocation();
    useEffect(() => {
        const selector = document.querySelector('ul.horizontal-menu a[href="' + window.location.pathname + '"]');
        if (selector) {
            selector.classList.add('active');
            const all: any = document.querySelectorAll('ul.horizontal-menu .nav-link.active');
            for (let i = 0; i < all.length; i++) {
                all[0]?.classList.remove('active');
            }
            const ul: any = selector.closest('ul.sub-menu');
            if (ul) {
                let ele: any = ul.closest('li.menu').querySelectorAll('.nav-link');
                if (ele) {
                    ele = ele[0];
                    setTimeout(() => {
                        ele?.classList.add('active');
                    });
                }
            }
        }
    }, [location]);

    const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;

    const themeConfig = useSelector((state: IRootState) => state.themeConfig);
    const [, setTheme] = useState<any>();
    const dispatch = useDispatch();

    const navigate = useNavigate();
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

    const submitForm = () => {
        navigate('/');
    };

    const [tabs, setTabs] = useState<string>('home');
    const toggleTabs = (name: string) => {
        setTabs(name);
    };

    const [setupmodal, setupModal] = useState(false);    
    const [qrmodal, qrModal] = useState(false);
    const [confmodal, confModal] = useState(false);
    const [signupmodal, signupModal] = useState(false);
    const [wizardmodal, wizardModal] = useState(false);

    const openQRModal = () => {
        setupModal(false); 
        qrModal(true); 
    };

    const openConfModal = () => {
        qrModal(false); 
        confModal(true); 
    };

    const { t } = useTranslation();

    const logoImage = (themeConfig.theme === 'light' || themeConfig.theme === 'system') ? (
            <img className="w-24 ltr:-ml-1 rtl:-mr-1 inline" src="/assets/images/logo/logo_dark.svg" alt="logo" />
      ) : (
        <img className="w-24 ltr:-ml-1 rtl:-mr-1 inline" src="/assets/images/logo/logo_light.svg" alt="logo" />
      );

    return (
        <header className={themeConfig.semidark && themeConfig.menu === 'horizontal' ? 'dark' : ''}>
            <div className="shadow-sm">
                <div className="relative bg-white flex w-full items-center px-5 py-2.5 dark:bg-black">
                    <div className="horizontal-logo flex lg:hidden justify-between items-center ltr:mr-2 rtl:ml-2">
                        <Link to="/" className="main-logo flex items-center shrink-0">
                            {logoImage}
                        </Link>
                        
                        <button
                            type="button"
                            className="collapse-icon flex-none dark:text-[#d0d2d6] hover:text-primary dark:hover:text-primary flex lg:hidden ltr:ml-2 rtl:mr-2 p-2 rounded-full bg-white-light/40 dark:bg-dark/40 hover:bg-white-light/90 dark:hover:bg-dark/60"
                            onClick={() => {
                                dispatch(toggleSidebar());
                            }}
                        >
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M20 7L4 7" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                <path opacity="0.5" d="M20 12L4 12" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                <path d="M20 17L4 17" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                            </svg>
                        </button>
                    </div>

                    <div className="sm:flex-1 ltr:sm:ml-0 ltr:ml-auto sm:rtl:mr-0 rtl:mr-auto flex items-center space-x-1.5 lg:space-x-2 rtl:space-x-reverse dark:text-[#d0d2d6]">
                         <div className="sm:ltr:mr-auto sm:rtl:ml-auto"></div> 
                        <div>
                            {themeConfig.theme === 'light' ? (
                                <button
                                    className={`${
                                        themeConfig.theme === 'light' &&
                                        'flex items-center p-2 rounded-full bg-white-light/40 dark:bg-dark/40 hover:text-primary hover:bg-white-light/90 dark:hover:bg-dark/60'
                                    }`}
                                    onClick={() => {
                                        setTheme('dark');
                                        dispatch(toggleTheme('dark'));
                                    }}
                                >
                                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <circle cx="12" cy="12" r="5" stroke="currentColor" strokeWidth="1.5" />
                                        <path d="M12 2V4" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path d="M12 20V22" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path d="M4 12L2 12" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path d="M22 12L20 12" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M19.7778 4.22266L17.5558 6.25424" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M4.22217 4.22266L6.44418 6.25424" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M6.44434 17.5557L4.22211 19.7779" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M19.7778 19.7773L17.5558 17.5551" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                    </svg>
                                </button>
                            ) : (
                                ''
                            )}
                            {themeConfig.theme === 'dark' && (
                                <button
                                    className={`${
                                        themeConfig.theme === 'dark' &&
                                        'flex items-center p-2 rounded-full bg-white-light/40 dark:bg-dark/40 hover:text-primary hover:bg-white-light/90 dark:hover:bg-dark/60'
                                    }`}
                                    onClick={() => {
                                        setTheme('system');
                                        dispatch(toggleTheme('system'));
                                    }}
                                >
                                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path
                                            d="M21.0672 11.8568L20.4253 11.469L21.0672 11.8568ZM12.1432 2.93276L11.7553 2.29085V2.29085L12.1432 2.93276ZM21.25 12C21.25 17.1086 17.1086 21.25 12 21.25V22.75C17.9371 22.75 22.75 17.9371 22.75 12H21.25ZM12 21.25C6.89137 21.25 2.75 17.1086 2.75 12H1.25C1.25 17.9371 6.06294 22.75 12 22.75V21.25ZM2.75 12C2.75 6.89137 6.89137 2.75 12 2.75V1.25C6.06294 1.25 1.25 6.06294 1.25 12H2.75ZM15.5 14.25C12.3244 14.25 9.75 11.6756 9.75 8.5H8.25C8.25 12.5041 11.4959 15.75 15.5 15.75V14.25ZM20.4253 11.469C19.4172 13.1373 17.5882 14.25 15.5 14.25V15.75C18.1349 15.75 20.4407 14.3439 21.7092 12.2447L20.4253 11.469ZM9.75 8.5C9.75 6.41182 10.8627 4.5828 12.531 3.57467L11.7553 2.29085C9.65609 3.5593 8.25 5.86509 8.25 8.5H9.75ZM12 2.75C11.9115 2.75 11.8077 2.71008 11.7324 2.63168C11.6686 2.56527 11.6538 2.50244 11.6503 2.47703C11.6461 2.44587 11.6482 2.35557 11.7553 2.29085L12.531 3.57467C13.0342 3.27065 13.196 2.71398 13.1368 2.27627C13.0754 1.82126 12.7166 1.25 12 1.25V2.75ZM21.7092 12.2447C21.6444 12.3518 21.5541 12.3539 21.523 12.3497C21.4976 12.3462 21.4347 12.3314 21.3683 12.2676C21.2899 12.1923 21.25 12.0885 21.25 12H22.75C22.75 11.2834 22.1787 10.9246 21.7237 10.8632C21.286 10.804 20.7293 10.9658 20.4253 11.469L21.7092 12.2447Z"
                                            fill="currentColor"
                                        />
                                    </svg>
                                </button>
                            )}
                            {themeConfig.theme === 'system' && (
                                <button
                                    className={`${
                                        themeConfig.theme === 'system' &&
                                        'flex items-center p-2 rounded-full bg-white-light/40 dark:bg-dark/40 hover:text-primary hover:bg-white-light/90 dark:hover:bg-dark/60'
                                    }`}
                                    onClick={() => {
                                        setTheme('light');
                                        dispatch(toggleTheme('light'));
                                    }}>
                                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path
                                            d="M3 9C3 6.17157 3 4.75736 3.87868 3.87868C4.75736 3 6.17157 3 9 3H15C17.8284 3 19.2426 3 20.1213 3.87868C21 4.75736 21 6.17157 21 9V14C21 15.8856 21 16.8284 20.4142 17.4142C19.8284 18 18.8856 18 17 18H7C5.11438 18 4.17157 18 3.58579 17.4142C3 16.8284 3 15.8856 3 14V9Z"
                                            stroke="currentColor"
                                            strokeWidth="1.5"
                                        />
                                        <path opacity="0.5" d="M22 21H2" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M15 15H9" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                    </svg>
                                </button>
                            )}
                        </div>

                        <div className="dropdown shrink-0">
                            <Dropdown
                                offset={[0, 8]}
                                placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                btnClassName="relative block p-2 rounded-full bg-white-light/40 dark:bg-dark/40 hover:text-primary hover:bg-white-light/90 dark:hover:bg-dark/60"
                                button={
                                    <span>                                    
                                        <svg className="ltr:mr-1 rtl:ml-1" width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <circle cx="12" cy="12" r="3" stroke="currentColor" stroke-width="1.5"></circle>
                                            <path opacity="0.5" d="M13.7654 2.15224C13.3978 2 12.9319 2 12 2C11.0681 2 10.6022 2 10.2346 2.15224C9.74457 2.35523 9.35522 2.74458 9.15223 3.23463C9.05957 3.45834 9.0233 3.7185 9.00911 4.09799C8.98826 4.65568 8.70226 5.17189 8.21894 5.45093C7.73564 5.72996 7.14559 5.71954 6.65219 5.45876C6.31645 5.2813 6.07301 5.18262 5.83294 5.15102C5.30704 5.08178 4.77518 5.22429 4.35436 5.5472C4.03874 5.78938 3.80577 6.1929 3.33983 6.99993C2.87389 7.80697 2.64092 8.21048 2.58899 8.60491C2.51976 9.1308 2.66227 9.66266 2.98518 10.0835C3.13256 10.2756 3.3397 10.437 3.66119 10.639C4.1338 10.936 4.43789 11.4419 4.43786 12C4.43783 12.5581 4.13375 13.0639 3.66118 13.3608C3.33965 13.5629 3.13248 13.7244 2.98508 13.9165C2.66217 14.3373 2.51966 14.8691 2.5889 15.395C2.64082 15.7894 2.87379 16.193 3.33973 17C3.80568 17.807 4.03865 18.2106 4.35426 18.4527C4.77508 18.7756 5.30694 18.9181 5.83284 18.8489C6.07289 18.8173 6.31632 18.7186 6.65204 18.5412C7.14547 18.2804 7.73556 18.27 8.2189 18.549C8.70224 18.8281 8.98826 19.3443 9.00911 19.9021C9.02331 20.2815 9.05957 20.5417 9.15223 20.7654C9.35522 21.2554 9.74457 21.6448 10.2346 21.8478C10.6022 22 11.0681 22 12 22C12.9319 22 13.3978 22 13.7654 21.8478C14.2554 21.6448 14.6448 21.2554 14.8477 20.7654C14.9404 20.5417 14.9767 20.2815 14.9909 19.902C15.0117 19.3443 15.2977 18.8281 15.781 18.549C16.2643 18.2699 16.8544 18.2804 17.3479 18.5412C17.6836 18.7186 17.927 18.8172 18.167 18.8488C18.6929 18.9181 19.2248 18.7756 19.6456 18.4527C19.9612 18.2105 20.1942 17.807 20.6601 16.9999C21.1261 16.1929 21.3591 15.7894 21.411 15.395C21.4802 14.8691 21.3377 14.3372 21.0148 13.9164C20.8674 13.7243 20.6602 13.5628 20.3387 13.3608C19.8662 13.0639 19.5621 12.558 19.5621 11.9999C19.5621 11.4418 19.8662 10.9361 20.3387 10.6392C20.6603 10.4371 20.8675 10.2757 21.0149 10.0835C21.3378 9.66273 21.4803 9.13087 21.4111 8.60497C21.3592 8.21055 21.1262 7.80703 20.6602 7C20.1943 6.19297 19.9613 5.78945 19.6457 5.54727C19.2249 5.22436 18.693 5.08185 18.1671 5.15109C17.9271 5.18269 17.6837 5.28136 17.3479 5.4588C16.8545 5.71959 16.2644 5.73002 15.7811 5.45096C15.2977 5.17191 15.0117 4.65566 14.9909 4.09794C14.9767 3.71848 14.9404 3.45833 14.8477 3.23463C14.6448 2.74458 14.2554 2.35523 13.7654 2.15224Z" stroke="currentColor" stroke-width="1.5">
                                            </path>                                            
                                        </svg>                                  
                                    </span>
                                }>
                                <ul className="!py-0 text-dark dark:text-white-dark w-[300px] sm:w-[350px] divide-y dark:divide-white/10">
                                    <div className="px-4 py-2 font-semibold">
                                        <h4 className="text-lg dark:text-white-light">Settings</h4>
                                    </div>  
                                    {/* 2FA Before verification */}
                                    <li>
                                        <div className='flex items-center py-3 px-5'>
                                            <div>
                                                <span className="grid place-content-center w-9 h-9 rounded-full bg-danger-light dark:bg-danger text-danger dark:text-danger-light">
                                                    <svg xmlns="http://www.w3.org/2000/svg" className="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                                                        <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path>
                                                    </svg>
                                                </span>
                                            </div>
                                            <span className="px-3 dark:text-gray-500">
                                                <div className="font-semibold text-sm dark:text-white-light/90 hover:cursor-pointer hover:underline underline-offset-4 decoration-dotted" onClick={() => setupModal(true)}>Two Factor Authentication</div>
                                                <div>To secure your account.</div>
                                            </span>
                                            <span className="font-semibold bg-danger-light dark:bg-danger rounded text-danger px-1 ltr:ml-auto rtl:mr-auto whitespace-pre dark:text-white/80 ltr:mr-2 rtl:ml-2">
                                                Required                                          
                                            </span>
                                        </div>
                                    </li>

                                    {/* 2FA After verification */}
                                    <li>
                                        <div className='flex items-center py-3 px-5'>
                                            <div>
                                                <span className="grid place-content-center w-9 h-9 rounded-full bg-success-light dark:bg-success text-success dark:text-success-light">
                                                    <svg xmlns="http://www.w3.org/2000/svg" className="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                                                        <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path>
                                                    </svg>
                                                </span>
                                            </div>
                                            <span className="px-3 dark:text-gray-500">
                                                <div className="font-semibold text-sm dark:text-white-light/90 hover:cursor-pointer hover:underline underline-offset-4 decoration-dotted" onClick={() => setupModal(true)}>Two Factor Authentication</div>
                                                <div>Successfully done</div>
                                            </span>
                                            <span className="font-semibold bg-success-light dark:bg-success rounded text-success px-1 ltr:ml-auto rtl:mr-auto whitespace-pre dark:text-white/80 ltr:mr-2 rtl:ml-2">
                                                Verified                                          
                                            </span>
                                        </div>
                                    </li>
                                   
                                   {/* Sign up Before completion */}
                                    <li>
                                        <div className='flex items-center py-3 px-5'>
                                            <div>
                                                <span className="grid place-content-center w-9 h-9 rounded-full bg-warning-light dark:bg-warning text-warning dark:text-warning-light">
                                                    <svg xmlns="http://www.w3.org/2000/svg" className="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">    
                                                        <circle cx="12" cy="12" r="10"></circle> 
                                                           <line x1="12" y1="8" x2="12" y2="12"></line>    
                                                           <line x1="12" y1="16" x2="12.01" y2="16"></line>
                                                    </svg>
                                                </span>
                                            </div>
                                            <span className="px-3 dark:text-gray-500">
                                                <div className="font-semibold text-sm dark:text-white-light/90 hover:cursor-pointer hover:underline underline-offset-4 decoration-dotted" onClick={() => signupModal(true)}>Signup</div>
                                                <div>Please complete signup process</div>
                                            </span>
                                            <span className="font-semibold bg-warning-light dark:bg-warning rounded text-warning px-1 ltr:ml-auto rtl:mr-auto whitespace-pre dark:text-white/80 ltr:mr-2 rtl:ml-2">
                                                20%                                          
                                            </span>
                                        </div>
                                    </li>

                                    {/* Sign up After completion */}
                                    <li>
                                        <div className='flex items-center py-3 px-5'>
                                            <div>
                                                <span className="grid place-content-center w-9 h-9 rounded-full bg-success-light dark:bg-success text-success dark:text-warning-light">
                                                    <svg xmlns="http://www.w3.org/2000/svg" className="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">    
                                                        <circle cx="12" cy="12" r="10"></circle> 
                                                           <line x1="12" y1="8" x2="12" y2="12"></line>    
                                                           <line x1="12" y1="16" x2="12.01" y2="16"></line>
                                                    </svg>
                                                </span>
                                            </div>
                                            <span className="px-3 dark:text-gray-500">
                                                <div className="font-semibold text-sm dark:text-white-light/90 hover:cursor-pointer hover:underline underline-offset-4 decoration-dotted" onClick={() => signupModal(true)}>Signup</div>
                                                <div>Successfully completed</div>
                                            </span>
                                            <span className="font-semibold bg-success-light dark:bg-success rounded text-success px-1 ltr:ml-auto rtl:mr-auto whitespace-pre dark:text-white/80 ltr:mr-2 rtl:ml-2">
                                                100%                                          
                                            </span>
                                        </div>
                                    </li>

                                    {/* Wizard */}
                                    <li>
                                        <div className='flex items-center py-3 px-5'>
                                            <div>
                                                <span className="grid place-content-center w-9 h-9 rounded-full bg-success-light dark:bg-success text-success dark:text-warning-light">
                                                    <svg className="w-5 h-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                        <path d="M14.5 8.5C14.5 9.88071 13.3807 11 12 11C10.6193 11 9.5 9.88071 9.5 8.5C9.5 7.11929 10.6193 6 12 6C13.3807 6 14.5 7.11929 14.5 8.5Z" fill="#000000"/>
                                                        <path d="M15.5812 16H8.50626C8.09309 16 7.87415 15.5411 8.15916 15.242C9.00598 14.3533 10.5593 13 12.1667 13C13.7899 13 15.2046 14.3801 15.947 15.2681C16.2011 15.5721 15.9774 16 15.5812 16Z" fill="#000000" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                                                        <circle cx="12" cy="12" r="10" stroke="#000000" stroke-width="2"/>
                                                    </svg>
                                                </span>
                                            </div>
                                            <span className="px-3 dark:text-gray-500">
                                                <div className="font-semibold  dark:text-white-light/90 hover:cursor-pointer hover:underline underline-offset-4 decoration-dotted" onClick={() => wizardModal(true)}>Wizard</div>
                                            </span>                                           
                                        </div>
                                    </li>
                                </ul>
                            </Dropdown>

                            {/* SETUP MODAL START */}
                            <Transition appear show={setupmodal} as={Fragment}>
                                <Dialog as="div" open={setupmodal} onClose={() => setupModal(false)}>
                                    <Transition.Child
                                        as={Fragment}
                                        enter="ease-out duration-300"
                                        enterFrom="opacity-0"
                                        enterTo="opacity-100"
                                        leave="ease-in duration-200"
                                        leaveFrom="opacity-100"
                                        leaveTo="opacity-0"
                                    >
                                        <div className="fixed inset-0" />
                                    </Transition.Child>
                                    <div className="fixed inset-0 z-[999] overflow-y-auto bg-[black]/60">
                                        <div className="flex min-h-screen items-center justify-center px-4">
                                            <Transition.Child
                                                as={Fragment}
                                                enter="ease-out duration-300"
                                                enterFrom="opacity-0 scale-95"
                                                enterTo="opacity-100 scale-100"
                                                leave="ease-in duration-200"
                                                leaveFrom="opacity-100 scale-100"
                                                leaveTo="opacity-0 scale-95"
                                            >
                                                <Dialog.Panel as="div" className="panel my-8 w-full max-w-lg overflow-hidden rounded-lg border-0 p-0 text-black dark:text-white-dark">
                                                    <div className="flex items-center justify-between bg-[#fbfbfb] px-5 py-3 dark:bg-[#121c2c]">
                                                        <h5 className="text-lg font-bold m-2">Setup Two Factor Authentication</h5>
                                                        <button type="button" className="text-white-dark hover:text-dark" onClick={() => setupModal(false)}>
                                                            <svg
                                                                xmlns="http://www.w3.org/2000/svg"
                                                                width="20"
                                                                height="20"
                                                                viewBox="0 0 24 24"
                                                                fill="none"
                                                                stroke="#282828"
                                                                strokeWidth="2.5"
                                                                strokeLinecap="round"
                                                                strokeLinejoin="round"
                                                            >
                                                                <line x1="18" y1="6" x2="6" y2="18"></line>
                                                                <line x1="6" y1="6" x2="18" y2="18"></line>
                                                            </svg>
                                                        </button>
                                                    </div>                                                    
                                                    <div className="p-6">
                                                    <p className="m-2 text-base"> 2FA provides an extra protection for your account by requiring a special code. Protect your account in just two steps:</p>

                                                    <div className="mx-6 my-3  text-base">
                                                        <ol className='list-decimal'>
                                                            <li className='pb-2'>Link a supported authentication app (such as Authenticator , Google Authenticator etc.)</li>
                                                            <li className='pb-2'>Enter the Confirmation Code</li>
                                                        </ol>
                                                    </div>

                                                    <p className="m-2 text-sm"> <span className='font-semibold'>Note:</span> You are only activating Two Factor Authentication for Owner Account only.</p>

                                                    <div className="mt-8 flex items-center justify-center">                                                        
                                                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]" onClick={openQRModal}>
                                                            NEXT
                                                        </button>
                                                    </div>
                                                    </div>
                                                </Dialog.Panel>
                                            </Transition.Child>
                                        </div>
                                    </div>
                                </Dialog>
                            </Transition>
                            {/* SETUP MODAL END */}

                            {/* QR MODAL START */}
                            <Transition appear show={qrmodal} as={Fragment}>
                                <Dialog as="div" open={qrmodal} onClose={() => qrModal(false)}>
                                    <Transition.Child
                                    as={Fragment}
                                    enter="ease-out duration-300"
                                    enterFrom="opacity-0"
                                    enterTo="opacity-100"
                                    leave="ease-in duration-200"
                                    leaveFrom="opacity-100"
                                    leaveTo="opacity-0"
                                    >
                                    <div className="fixed inset-0" />
                                    </Transition.Child>
                                    <div className="fixed inset-0 z-[999] overflow-y-auto bg-[black]/60">
                                    <div className="flex min-h-screen items-center justify-center px-4">
                                        <Transition.Child
                                        as={Fragment}
                                        enter="ease-out duration-300"
                                        enterFrom="opacity-0 scale-95"
                                        enterTo="opacity-100 scale-100"
                                        leave="ease-in duration-200"
                                        leaveFrom="opacity-100 scale-100"
                                        leaveTo="opacity-0 scale-95"
                                        >
                                        <Dialog.Panel as="div" className="panel my-8 w-full max-w-lg overflow-hidden rounded-lg border-0 p-0 text-black dark:text-white-dark">
                                            <div className="flex items-center justify-between bg-[#fbfbfb] px-5 py-3 dark:bg-[#121c2c]">
                                            <h5 className="text-lg font-bold m-2">Two Factor Authentication</h5>
                                            <button type="button" className="text-white-dark hover:text-dark" onClick={() => qrModal(false)}>
                                                <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                width="20"
                                                height="20"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke="#282828"
                                                strokeWidth="2.5"
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                >
                                                <line x1="18" y1="6" x2="6" y2="18"></line>
                                                <line x1="6" y1="6" x2="18" y2="18"></line>
                                                </svg>
                                            </button>
                                            </div>
                                            <div className="p-6">
                                                <p className="m-2 text-base"> Scan the QR code using supported authenticator app</p>
                                                <div className="flex justify-center">
                                                    <div><img className="h-40" src="/assets/images/logo/qr-code.svg" alt="qr code img" /></div>                   
                                                </div>
                                                <p className="m-2 text-base"> Can't scan the QR code?</p>
                                                <p className="m-2 text-base"> Copy this "Recovery codes" to your authenticator app. Enter this code into your authenticator app instead. </p>
                                                
                                                <div className="mt-6 text-base flex justify-between font-semibold">
                                                    <div><p className='text-black'>Recovery Codes:</p></div>
                                                    <div>
                                                        <p className='text-orange-600 flex'>
                                                            <Link to="/" className=' flex'>
                                                                Download <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className='mx-1 mt-1 font-bold'>
                                                                    <path d="M12.5535 16.5061C12.4114 16.6615 12.2106 16.75 12 16.75C11.7894 16.75 11.5886 16.6615 11.4465 16.5061L7.44648 12.1311C7.16698 11.8254 7.18822 11.351 7.49392 11.0715C7.79963 10.792 8.27402 10.8132 8.55352 11.1189L11.25 14.0682V3C11.25 2.58579 11.5858 2.25 12 2.25C12.4142 2.25 12.75 2.58579 12.75 3V14.0682L15.4465 11.1189C15.726 10.8132 16.2004 10.792 16.5061 11.0715C16.8118 11.351 16.833 11.8254 16.5535 12.1311L12.5535 16.5061Z" fill="#c8400d"/>
                                                                    <path d="M3.75 15C3.75 14.5858 3.41422 14.25 3 14.25C2.58579 14.25 2.25 14.5858 2.25 15V15.0549C2.24998 16.4225 2.24996 17.5248 2.36652 18.3918C2.48754 19.2919 2.74643 20.0497 3.34835 20.6516C3.95027 21.2536 4.70814 21.5125 5.60825 21.6335C6.47522 21.75 7.57754 21.75 8.94513 21.75H15.0549C16.4225 21.75 17.5248 21.75 18.3918 21.6335C19.2919 21.5125 20.0497 21.2536 20.6517 20.6516C21.2536 20.0497 21.5125 19.2919 21.6335 18.3918C21.75 17.5248 21.75 16.4225 21.75 15.0549V15C21.75 14.5858 21.4142 14.25 21 14.25C20.5858 14.25 20.25 14.5858 20.25 15C20.25 16.4354 20.2484 17.4365 20.1469 18.1919C20.0482 18.9257 19.8678 19.3142 19.591 19.591C19.3142 19.8678 18.9257 20.0482 18.1919 20.1469C17.4365 20.2484 16.4354 20.25 15 20.25H9C7.56459 20.25 6.56347 20.2484 5.80812 20.1469C5.07435 20.0482 4.68577 19.8678 4.40901 19.591C4.13225 19.3142 3.9518 18.9257 3.85315 18.1919C3.75159 17.4365 3.75 16.4354 3.75 15Z" fill="#c8400d"/>
                                                                    </svg>
                                                            </Link>
                                                            <span className='px-3 text-gray-500'> | </span> 
                                                            <Link to="/"  className=' flex'>
                                                                Copy <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className='mx-1 mt-1 font-bold'>
                                                                <path fill-rule="evenodd" clip-rule="evenodd" d="M19.5 16.5L19.5 4.5L18.75 3.75H9L8.25 4.5L8.25 7.5L5.25 7.5L4.5 8.25V20.25L5.25 21H15L15.75 20.25V17.25H18.75L19.5 16.5ZM15.75 15.75L15.75 8.25L15 7.5L9.75 7.5V5.25L18 5.25V15.75H15.75ZM6 9L14.25 9L14.25 19.5L6 19.5L6 9Z" fill="#c8400d"/>
                                                                </svg>
                                                            </Link>                            
                                                        </p>
                                                    </div>
                                                </div>

                                                <pre>
                                                    <div className='border border-gray-500 p-3 my-3 grid grid-cols-3 gap-2 text-center'>                    
                                                        <div>123456</div>
                                                        <div>456789</div>
                                                        <div>789123</div>
                                                        <div>147258</div>
                                                        <div>258369</div>
                                                        <div>369147</div>
                                                        <div>159236</div>
                                                        <div>159478</div>
                                                        <div>784512</div>
                                                        <div>895623</div>
                                                        <div>741963</div>
                                                        <div>852741</div>
                                                    </div>
                                                </pre>

                                                <div className="flex justify-center py-6">                                                    
                                                    <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]" onClick={openConfModal}>
                                                        ENTER CONFIRMATION CODE
                                                    </button>                                                    
                                                </div>  
                                            </div>
                                        </Dialog.Panel>
                                        </Transition.Child>
                                    </div>
                                    </div>
                                </Dialog>
                            </Transition>
                            {/* QR MODAL END */}

                            {/* CONFIRMATION MODAL START */}
                            <Transition appear show={confmodal} as={Fragment}>
                                <Dialog as="div" open={confmodal} onClose={() => confModal(false)}>
                                    <Transition.Child
                                    as={Fragment}
                                    enter="ease-out duration-300"
                                    enterFrom="opacity-0"
                                    enterTo="opacity-100"
                                    leave="ease-in duration-200"
                                    leaveFrom="opacity-100"
                                    leaveTo="opacity-0"
                                    >
                                    <div className="fixed inset-0" />
                                    </Transition.Child>
                                    <div className="fixed inset-0 z-[999] overflow-y-auto bg-[black]/60">
                                        <div className="flex min-h-screen items-center justify-center px-4">
                                            <Transition.Child
                                            as={Fragment}
                                            enter="ease-out duration-300"
                                            enterFrom="opacity-0 scale-95"
                                            enterTo="opacity-100 scale-100"
                                            leave="ease-in duration-200"
                                            leaveFrom="opacity-100 scale-100"
                                            leaveTo="opacity-0 scale-95"
                                            >
                                            <Dialog.Panel as="div" className="panel my-8 w-full max-w-lg overflow-hidden rounded-lg border-0 p-0 text-black dark:text-white-dark">
                                                <div className="flex items-center justify-between bg-[#fbfbfb] px-5 py-3 dark:bg-[#121c2c]">
                                                <h5 className="text-lg font-bold m-2">Confirmation Two Factor Authentication</h5>
                                                <button type="button" className="text-white-dark hover:text-dark" onClick={() => confModal(false)}>
                                                    <svg
                                                    xmlns="http://www.w3.org/2000/svg"
                                                    width="20"
                                                    height="20"
                                                    viewBox="0 0 24 24"
                                                    fill="none"
                                                    stroke="#282828"
                                                    strokeWidth="2.5"
                                                    strokeLinecap="round"
                                                    strokeLinejoin="round"
                                                    >
                                                    <line x1="18" y1="6" x2="6" y2="18"></line>
                                                    <line x1="6" y1="6" x2="18" y2="18"></line>
                                                    </svg>
                                                </button>
                                                </div>
                                                <div className="p-6">
                                                    <p className="m-2 text-base">Enter 6-digit code generated by your authenticator app to activate your 2FA</p>
                                                    
                                                    <div className='grid grid-cols-6 gap-4 m-8'>
                                                        <div>
                                                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400 text-center" />
                                                        </div>
                                                        <div>
                                                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400 text-center" />
                                                        </div>
                                                        <div>
                                                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400 text-center" />
                                                        </div>
                                                        <div>
                                                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400 text-center" />
                                                        </div>
                                                        <div>
                                                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400 text-center" />
                                                        </div>
                                                        <div>
                                                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400 text-center" />
                                                        </div>
                                                    </div>
                                                        
                                                    <div className="flex justify-center py-6">
                                                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]" onClick={() => coloredToast('dark')}>
                                                            CONFIRM
                                                        </button>
                                                    </div>      
                                                </div>
                                            </Dialog.Panel>
                                            </Transition.Child>
                                        </div>
                                    </div>
                                </Dialog>
                            </Transition>
                            {/* CONFIRMATION MODAL END */}

                            {/* SIGNUP MODAL START */}
                            <Transition appear show={signupmodal} as={Fragment}>
                                <Dialog as="div" open={signupmodal} onClose={() => signupModal(false)}>
                                    <Transition.Child
                                        as={Fragment}
                                        enter="ease-out duration-300"
                                        enterFrom="opacity-0"
                                        enterTo="opacity-100"
                                        leave="ease-in duration-200"
                                        leaveFrom="opacity-100"
                                        leaveTo="opacity-0"
                                    >
                                        <div className="fixed inset-0" />
                                    </Transition.Child>
                                    <div className="fixed inset-0 z-[999] overflow-y-auto bg-[black]/60">
                                        <div className="flex min-h-screen items-center justify-center px-4">
                                            <Transition.Child
                                                as={Fragment}
                                                enter="ease-out duration-300"
                                                enterFrom="opacity-0 scale-95"
                                                enterTo="opacity-100 scale-100"
                                                leave="ease-in duration-200"
                                                leaveFrom="opacity-100 scale-100"
                                                leaveTo="opacity-0 scale-95"
                                            >
                                                <Dialog.Panel as="div" className="panel my-8 w-full max-w-lg overflow-hidden rounded-lg border-0 p-0 text-black dark:text-white-dark">
                                                    <div className="flex items-center justify-between bg-[#fbfbfb] px-5 py-3 dark:bg-[#121c2c]">
                                                        <h5 className="text-lg font-bold m-2">To Complete SignUp</h5>
                                                        <button type="button" className="text-white-dark hover:text-dark" onClick={() => signupModal(false)}>
                                                            <svg
                                                                xmlns="http://www.w3.org/2000/svg"
                                                                width="20"
                                                                height="20"
                                                                viewBox="0 0 24 24"
                                                                fill="none"
                                                                stroke="#282828"
                                                                strokeWidth="2.5"
                                                                strokeLinecap="round"
                                                                strokeLinejoin="round"
                                                            >
                                                                <line x1="18" y1="6" x2="6" y2="18"></line>
                                                                <line x1="6" y1="6" x2="18" y2="18"></line>
                                                            </svg>
                                                        </button>
                                                    </div>                                                    
                                                    
                                                    <form className="space-y-5 p-8" onSubmit={submitForm}>
                                                        <p className="mb-7">Please fill all details to complete Registration</p>
                                                        
                                                        <div>
                                                            <label htmlFor="changepwd">Name <span className='text-red-600'>*</span></label>
                                                            <input id="phoneno" type="text" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Name" />
                                                        </div>

                                                        <div>
                                                            <label htmlFor="email">Email <span className='text-red-600'>*</span></label> 
                                                            <input id="email" type="email" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Email" />
                                                        </div>

                                                        <div>
                                                            <label htmlFor="password">Password <span className='text-red-600'>*</span></label>
                                                            <input id="password" type="password" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Password" />
                                                        </div>

                                                        <div>
                                                            <label htmlFor="changepwd">Date of Birth <span className='text-red-600'>*</span></label>
                                                            <input id="phoneno" type="date" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter DOB" />
                                                        </div>

                                                        <div>
                                                            <label htmlFor="changepwd">Account Type <span className='text-red-600'>*</span></label>
                                                            <select className="form-select text-white-dark border border-gray-400 focus:border-orange-400">
                                                                <option value="" selected disabled>Select</option>
                                                                <option value="owner">Owner</option>
                                                                <option value="client">Client</option>
                                                            </select>
                                                        </div>

                                                        <div className="flex justify-center py-6">
                                                        <Link to="/">
                                                            <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                                                                SIGN UP
                                                            </button>
                                                        </Link>
                                                        </div> 
                                                    </form>  
                                                </Dialog.Panel>
                                            </Transition.Child>
                                        </div>
                                    </div>
                                </Dialog>
                            </Transition>
                            {/* SIGNUP MODAL END */}

                            {/* WIZARD MODAL START */}
                                <Transition appear show={wizardmodal} as={Fragment}>
                                    <Dialog as="div" open={wizardmodal} onClose={() => wizardModal(false)}>
                                        <Transition.Child
                                            as={Fragment}
                                            enter="ease-out duration-300"
                                            enterFrom="opacity-0"
                                            enterTo="opacity-100"
                                            leave="ease-in duration-200"
                                            leaveFrom="opacity-100"
                                            leaveTo="opacity-0"
                                        >
                                            <div className="fixed inset-0" />
                                        </Transition.Child>
                                        <div className="fixed inset-0 z-[999] bg-[black]/60">
                                            <div className="flex min-h-screen items-start justify-center px-4">
                                                <Transition.Child
                                                    as={Fragment}
                                                    enter="ease-out duration-300"
                                                    enterFrom="opacity-0 scale-95"
                                                    enterTo="opacity-100 scale-100"
                                                    leave="ease-in duration-200"
                                                    leaveFrom="opacity-100 scale-100"
                                                    leaveTo="opacity-0 scale-95"
                                                >
                                                    <Dialog.Panel className="panel my-8 w-full max-w-5xl overflow-hidden rounded-lg border-0 p-0 text-black dark:text-white-dark">
                                                        <div className="flex items-center justify-between bg-[#fbfbfb] px-5 py-3 dark:bg-[#121c2c]">
                                                            <h5 className="text-lg font-bold">Widgets</h5>
                                                            <button onClick={() => wizardModal(false)} type="button" className="text-white-dark hover:text-dark">
                                                                <svg
                                                                    xmlns="http://www.w3.org/2000/svg"
                                                                    width="20"
                                                                    height="20"
                                                                    viewBox="0 0 24 24"
                                                                    fill="none"
                                                                    stroke="currentColor"
                                                                    strokeWidth="1.5"
                                                                    strokeLinecap="round"
                                                                    strokeLinejoin="round"
                                                                >
                                                                    <line x1="18" y1="6" x2="6" y2="18"></line>
                                                                    <line x1="6" y1="6" x2="18" y2="18"></line>
                                                                </svg>
                                                            </button>
                                                        </div>
                                                        <div className="p-5">
                                                            {/*  */}
                                                            <div>
                                                                <ul className="sm:flex font-semibold border-b border-[#ebedf2] dark:border-[#191e3a] mb-5 whitespace-nowrap overflow-y-auto">
                                                                    
                                                                    <li className="inline-block">
                                                                        <button
                                                                            onClick={() => toggleTabs('home')}
                                                                            className={`flex gap-2 p-4 border-b border-transparent hover:border-primary hover:text-primary ${tabs === 'home' ? '!border-primary text-primary' : ''}`}
                                                                        >
                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                <circle opacity="0.5" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="1.5" />
                                                                                <path d="M12 6V18" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                                                                <path
                                                                                    d="M15 9.5C15 8.11929 13.6569 7 12 7C10.3431 7 9 8.11929 9 9.5C9 10.8807 10.3431 12 12 12C13.6569 12 15 13.1193 15 14.5C15 15.8807 13.6569 17 12 17C10.3431 17 9 15.8807 9 14.5"
                                                                                    stroke="currentColor"
                                                                                    strokeWidth="1.5"
                                                                                    strokeLinecap="round"
                                                                                />
                                                                            </svg>
                                                                            User Profile
                                                                        </button>
                                                                    </li>
                                                                    <li className="inline-block">
                                                                        <button
                                                                            onClick={() => toggleTabs('payment-details')}
                                                                            className={`flex gap-2 p-4 border-b border-transparent hover:border-primary hover:text-primary ${tabs === 'payment-details' ? '!border-primary text-primary' : ''}`}
                                                                        >
                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                <circle opacity="0.5" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="1.5" />
                                                                                <path d="M12 6V18" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                                                                <path
                                                                                    d="M15 9.5C15 8.11929 13.6569 7 12 7C10.3431 7 9 8.11929 9 9.5C9 10.8807 10.3431 12 12 12C13.6569 12 15 13.1193 15 14.5C15 15.8807 13.6569 17 12 17C10.3431 17 9 15.8807 9 14.5"
                                                                                    stroke="currentColor"
                                                                                    strokeWidth="1.5"
                                                                                    strokeLinecap="round"
                                                                                />
                                                                            </svg>
                                                                            Payment Details
                                                                        </button>
                                                                    </li>
                                                                    <li className="inline-block">
                                                                        <button
                                                                            onClick={() => toggleTabs('info')}
                                                                            className={`flex gap-2 p-4 border-b border-transparent hover:border-primary hover:text-primary ${tabs === 'info' ? '!border-primary text-primary' : ''}`}
                                                                        >
                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                <path
                                                                                    opacity="0.5"
                                                                                    d="M2 12.2039C2 9.91549 2 8.77128 2.5192 7.82274C3.0384 6.87421 3.98695 6.28551 5.88403 5.10813L7.88403 3.86687C9.88939 2.62229 10.8921 2 12 2C13.1079 2 14.1106 2.62229 16.116 3.86687L18.116 5.10812C20.0131 6.28551 20.9616 6.87421 21.4808 7.82274C22 8.77128 22 9.91549 22 12.2039V13.725C22 17.6258 22 19.5763 20.8284 20.7881C19.6569 22 17.7712 22 14 22H10C6.22876 22 4.34315 22 3.17157 20.7881C2 19.5763 2 17.6258 2 13.725V12.2039Z"
                                                                                    stroke="currentColor"
                                                                                    strokeWidth="1.5"
                                                                                />
                                                                                <path d="M12 15L12 18" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                                                            </svg>
                                                                            General Information
                                                                        </button>
                                                                    </li>
                                                                </ul>
                                                            </div>
                                                            
                                                            {tabs === 'home' ? (
                                                                <div>
                                                                    <form className="border border-[#ebedf2] dark:border-[#191e3a] rounded-md p-4 mb-5 bg-white dark:bg-black">
                                                                        <h6 className="text-lg font-bold mb-5">Profile Information <span className='text-red-600'>*</span></h6>
                                                                        <div className="flex flex-col sm:flex-row">
                                                                            <div className="ltr:sm:mr-4 rtl:sm:ml-4 w-full sm:w-2/12 mb-5">
                                                                                <img src="/assets//images/profile-34.jpeg" alt="img" className="w-20 h-20 md:w-32 md:h-32 rounded-full object-cover mx-auto" />
                                                                            </div>
                                                                            <div className="flex-1 grid grid-cols-1 sm:grid-cols-2 gap-5">
                                                                                <div>
                                                                                    <label htmlFor="name">Full Name</label>
                                                                                    <input id="name" type="text" placeholder="Jimmy Turner" className="form-input" />
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="profession">Profession</label>
                                                                                    <input id="profession" type="text" placeholder="Web Developer" className="form-input" />
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="country">Country</label>
                                                                                    <select defaultValue="United States" id="country" className="form-select text-white-dark">
                                                                                        <option value="All Countries">All Countries</option>
                                                                                        <option value="United States">United States</option>
                                                                                        <option value="India">India</option>
                                                                                        <option value="Japan">Japan</option>
                                                                                        <option value="China">China</option>
                                                                                        <option value="Brazil">Brazil</option>
                                                                                        <option value="Norway">Norway</option>
                                                                                        <option value="Canada">Canada</option>
                                                                                    </select>
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="address">Address</label>
                                                                                    <input id="address" type="text" placeholder="New York" className="form-input" />
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="location">Location</label>
                                                                                    <input id="location" type="text" placeholder="Location" className="form-input" />
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="phone">Phone</label>
                                                                                    <input id="phone" type="text" placeholder="+1 (530) 555-12121" className="form-input" />
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="email">Email</label>
                                                                                    <input id="email" type="email" placeholder="Jimmy@gmail.com" className="form-input" />
                                                                                </div>
                                                                                <div>
                                                                                    <label htmlFor="web">Website</label>
                                                                                    <input id="web" type="text" placeholder="Enter URL" className="form-input" />
                                                                                </div>                                                                                
                                                                                <div className="sm:col-span-2 mt-3 flex justify-center">
                                                                                    <button type="button" className="btn btn-secondary rounded-full">
                                                                                        Save
                                                                                    </button>
                                                                                </div>
                                                                            </div>
                                                                        </div>
                                                                    </form>
                                                                   
                                                                </div>
                                                            ) : (
                                                                ''
                                                            )}
                                                            {tabs === 'payment-details' ? (
                                                                <div>
                                                                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-5 mb-5">
                                                                        <div className="panel">
                                                                            <div className="mb-5">
                                                                                <h5 className="font-semibold text-lg mb-4">Billing Address</h5>
                                                                                <p>
                                                                                    Changes to your <span className="text-primary">Billing</span> information will take effect starting with scheduled payment and will be refelected on your next
                                                                                    invoice.
                                                                                </p>
                                                                            </div>
                                                                            <div className="mb-5">
                                                                                <div className="border-b border-[#ebedf2] dark:border-[#1b2e4b]">
                                                                                    <div className="flex items-start justify-between py-3">
                                                                                        <h6 className="text-[#515365] font-bold dark:text-white-dark text-[15px]">
                                                                                            Address #1
                                                                                            <span className="block text-white-dark dark:text-white-light font-normal text-xs mt-1">2249 Caynor Circle, New Brunswick, New Jersey</span>
                                                                                        </h6>
                                                                                        <div className="flex items-start justify-between ltr:ml-auto rtl:mr-auto">
                                                                                            <button className="btn btn-dark">Edit</button>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                                <div className="border-b border-[#ebedf2] dark:border-[#1b2e4b]">
                                                                                    <div className="flex items-start justify-between py-3">
                                                                                        <h6 className="text-[#515365] font-bold dark:text-white-dark text-[15px]">
                                                                                            Address #2
                                                                                            <span className="block text-white-dark dark:text-white-light font-normal text-xs mt-1">4262 Leverton Cove Road, Springfield, Massachusetts</span>
                                                                                        </h6>
                                                                                        <div className="flex items-start justify-between ltr:ml-auto rtl:mr-auto">
                                                                                            <button className="btn btn-dark">Edit</button>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                                <div>
                                                                                    <div className="flex items-start justify-between py-3">
                                                                                        <h6 className="text-[#515365] font-bold dark:text-white-dark text-[15px]">
                                                                                            Address #3
                                                                                            <span className="block text-white-dark dark:text-white-light font-normal text-xs mt-1">2692 Berkshire Circle, Knoxville, Tennessee</span>
                                                                                        </h6>
                                                                                        <div className="flex items-start justify-between ltr:ml-auto rtl:mr-auto">
                                                                                            <button className="btn btn-dark">Edit</button>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                            </div>
                                                                            <button className="btn btn-primary">Add Address</button>
                                                                        </div>
                                                                        <div className="panel">
                                                                            <div className="mb-5">
                                                                                <h5 className="font-semibold text-lg mb-4">Payment History</h5>
                                                                                <p>
                                                                                    Changes to your <span className="text-primary">Payment Method</span> information will take effect starting with scheduled payment and will be refelected on your
                                                                                    next invoice.
                                                                                </p>
                                                                            </div>
                                                                            <div className="mb-5">
                                                                                <div className="border-b border-[#ebedf2] dark:border-[#1b2e4b]">
                                                                                    <div className="flex items-start justify-between py-3">
                                                                                        <div className="flex-none ltr:mr-4 rtl:ml-4">
                                                                                            <img src="/assets/images/card-americanexpress.svg" alt="img" />
                                                                                        </div>
                                                                                        <h6 className="text-[#515365] font-bold dark:text-white-dark text-[15px]">
                                                                                            Mastercard
                                                                                            <span className="block text-white-dark dark:text-white-light font-normal text-xs mt-1">XXXX XXXX XXXX 9704</span>
                                                                                        </h6>
                                                                                        <div className="flex items-start justify-between ltr:ml-auto rtl:mr-auto">
                                                                                            <button className="btn btn-dark">Edit</button>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                                <div className="border-b border-[#ebedf2] dark:border-[#1b2e4b]">
                                                                                    <div className="flex items-start justify-between py-3">
                                                                                        <div className="flex-none ltr:mr-4 rtl:ml-4">
                                                                                            <img src="/assets/images/card-mastercard.svg" alt="img" />
                                                                                        </div>
                                                                                        <h6 className="text-[#515365] font-bold dark:text-white-dark text-[15px]">
                                                                                            American Express
                                                                                            <span className="block text-white-dark dark:text-white-light font-normal text-xs mt-1">XXXX XXXX XXXX 310</span>
                                                                                        </h6>
                                                                                        <div className="flex items-start justify-between ltr:ml-auto rtl:mr-auto">
                                                                                            <button className="btn btn-dark">Edit</button>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                                <div>
                                                                                    <div className="flex items-start justify-between py-3">
                                                                                        <div className="flex-none ltr:mr-4 rtl:ml-4">
                                                                                            <img src="/assets/images/card-visa.svg" alt="img" />
                                                                                        </div>
                                                                                        <h6 className="text-[#515365] font-bold dark:text-white-dark text-[15px]">
                                                                                            Visa
                                                                                            <span className="block text-white-dark dark:text-white-light font-normal text-xs mt-1">XXXX XXXX XXXX 5264</span>
                                                                                        </h6>
                                                                                        <div className="flex items-start justify-between ltr:ml-auto rtl:mr-auto">
                                                                                            <button className="btn btn-dark">Edit</button>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                            </div>
                                                                            <button className="btn btn-primary">Add Payment Method</button>
                                                                        </div>
                                                                    </div>
                                                                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-5">
                                                                        <div className="panel">
                                                                            <div className="mb-5">
                                                                                <h5 className="font-semibold text-lg mb-4">Add Billing Address</h5>
                                                                                <p>
                                                                                    Changes your New <span className="text-primary">Billing</span> Information.
                                                                                </p>
                                                                            </div>
                                                                            <div className="mb-5">
                                                                                <form>
                                                                                    <div className="mb-5 grid grid-cols-1 sm:grid-cols-2 gap-4">
                                                                                        <div>
                                                                                            <label htmlFor="billingName">Name</label>
                                                                                            <input id="billingName" type="text" placeholder="Enter Name" className="form-input" />
                                                                                        </div>
                                                                                        <div>
                                                                                            <label htmlFor="billingEmail">Email</label>
                                                                                            <input id="billingEmail" type="email" placeholder="Enter Email" className="form-input" />
                                                                                        </div>
                                                                                    </div>
                                                                                    <div className="mb-5">
                                                                                        <label htmlFor="billingAddress">Address</label>
                                                                                        <input id="billingAddress" type="text" placeholder="Enter Address" className="form-input" />
                                                                                    </div>
                                                                                    <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4 mb-5">
                                                                                        <div className="md:col-span-2">
                                                                                            <label htmlFor="billingCity">City</label>
                                                                                            <input id="billingCity" type="text" placeholder="Enter City" className="form-input" />
                                                                                        </div>
                                                                                        <div>
                                                                                            <label htmlFor="billingState">State</label>
                                                                                            <select id="billingState" className="form-select text-white-dark">
                                                                                                <option>Choose...</option>
                                                                                                <option>...</option>
                                                                                            </select>
                                                                                        </div>
                                                                                        <div>
                                                                                            <label htmlFor="billingZip">Zip</label>
                                                                                            <input id="billingZip" type="text" placeholder="Enter Zip" className="form-input" />
                                                                                        </div>
                                                                                    </div>
                                                                                    <button type="button" className="btn btn-primary">
                                                                                        Add
                                                                                    </button>
                                                                                </form>
                                                                            </div>
                                                                        </div>
                                                                        <div className="panel">
                                                                            <div className="mb-5">
                                                                                <h5 className="font-semibold text-lg mb-4">Add Payment Method</h5>
                                                                                <p>
                                                                                    Changes your New <span className="text-primary">Payment Method </span>
                                                                                    Information.
                                                                                </p>
                                                                            </div>
                                                                            <div className="mb-5">
                                                                                <form>
                                                                                    <div className="mb-5 grid grid-cols-1 sm:grid-cols-2 gap-4">
                                                                                        <div>
                                                                                            <label htmlFor="payBrand">Card Brand</label>
                                                                                            <select id="payBrand" className="form-select text-white-dark">
                                                                                                <option value="Mastercard">Mastercard</option>
                                                                                                <option value="American Express">American Express</option>
                                                                                                <option value="Visa">Visa</option>
                                                                                                <option value="Discover">Discover</option>
                                                                                            </select>
                                                                                        </div>
                                                                                        <div>
                                                                                            <label htmlFor="payNumber">Card Number</label>
                                                                                            <input id="payNumber" type="text" placeholder="Card Number" className="form-input" />
                                                                                        </div>
                                                                                    </div>
                                                                                    <div className="mb-5 grid grid-cols-1 sm:grid-cols-2 gap-4">
                                                                                        <div>
                                                                                            <label htmlFor="payHolder">Holder Name</label>
                                                                                            <input id="payHolder" type="text" placeholder="Holder Name" className="form-input" />
                                                                                        </div>
                                                                                        <div>
                                                                                            <label htmlFor="payCvv">CVV/CVV2</label>
                                                                                            <input id="payCvv" type="text" placeholder="CVV" className="form-input" />
                                                                                        </div>
                                                                                    </div>
                                                                                    <div className="mb-5 grid grid-cols-1 sm:grid-cols-2 gap-4">
                                                                                        <div>
                                                                                            <label htmlFor="payExp">Card Expiry</label>
                                                                                            <input id="payExp" type="text" placeholder="Card Expiry" className="form-input" />
                                                                                        </div>
                                                                                    </div>
                                                                                    <button type="button" className="btn btn-primary">
                                                                                        Add
                                                                                    </button>
                                                                                </form>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            ) : (
                                                                ''
                                                            )}  

                                                            {tabs === 'info' ? (
                                                                <div>
                                                                   <div className="pt-5">
                                                                        <div className="grid grid-cols-1 lg:grid-cols-3 gap-5 mb-5">
                                                                            {/* profile details */}
                                                                            <div className="panel">
                                                                                <div className="mb-5">
                                                                                    <div className="flex flex-col justify-center items-center">
                                                                                        <img src="/assets/images/profile-34.jpeg" alt="img" className="w-24 h-24 rounded-full object-cover  mb-5" />
                                                                                        <p className="font-semibold text-primary text-xl">Jimmy Turner</p>
                                                                                    </div>
                                                                                    <ul className="mt-5 flex flex-col max-w-[160px] m-auto space-y-4 font-semibold text-white-dark">
                                                                                        <li className="flex items-center gap-2">
                                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                                <path
                                                                                                    d="M2.3153 12.6978C2.26536 12.2706 2.2404 12.057 2.2509 11.8809C2.30599 10.9577 2.98677 10.1928 3.89725 10.0309C4.07094 10 4.286 10 4.71612 10H15.2838C15.7139 10 15.929 10 16.1027 10.0309C17.0132 10.1928 17.694 10.9577 17.749 11.8809C17.7595 12.057 17.7346 12.2706 17.6846 12.6978L17.284 16.1258C17.1031 17.6729 16.2764 19.0714 15.0081 19.9757C14.0736 20.6419 12.9546 21 11.8069 21H8.19303C7.04537 21 5.9263 20.6419 4.99182 19.9757C3.72352 19.0714 2.89681 17.6729 2.71598 16.1258L2.3153 12.6978Z"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                />
                                                                                                <path opacity="0.5" d="M17 17H19C20.6569 17 22 15.6569 22 14C22 12.3431 20.6569 11 19 11H17.5" stroke="currentColor" strokeWidth="1.5" />
                                                                                                <path
                                                                                                    opacity="0.5"
                                                                                                    d="M10.0002 2C9.44787 2.55228 9.44787 3.44772 10.0002 4C10.5524 4.55228 10.5524 5.44772 10.0002 6"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                    strokeLinecap="round"
                                                                                                    strokeLinejoin="round"
                                                                                                />
                                                                                                <path
                                                                                                    d="M4.99994 7.5L5.11605 7.38388C5.62322 6.87671 5.68028 6.0738 5.24994 5.5C4.81959 4.9262 4.87665 4.12329 5.38382 3.61612L5.49994 3.5"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                    strokeLinecap="round"
                                                                                                    strokeLinejoin="round"
                                                                                                />
                                                                                                <path
                                                                                                    d="M14.4999 7.5L14.6161 7.38388C15.1232 6.87671 15.1803 6.0738 14.7499 5.5C14.3196 4.9262 14.3767 4.12329 14.8838 3.61612L14.9999 3.5"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                    strokeLinecap="round"
                                                                                                    strokeLinejoin="round"
                                                                                                />
                                                                                            </svg>{' '}
                                                                                            Web Developer
                                                                                        </li>
                                                                                        <li className="flex items-center gap-2">
                                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                                <path
                                                                                                    d="M2 12C2 8.22876 2 6.34315 3.17157 5.17157C4.34315 4 6.22876 4 10 4H14C17.7712 4 19.6569 4 20.8284 5.17157C22 6.34315 22 8.22876 22 12V14C22 17.7712 22 19.6569 20.8284 20.8284C19.6569 22 17.7712 22 14 22H10C6.22876 22 4.34315 22 3.17157 20.8284C2 19.6569 2 17.7712 2 14V12Z"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                />
                                                                                                <path opacity="0.5" d="M7 4V2.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                                                                                <path opacity="0.5" d="M17 4V2.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                                                                                <path opacity="0.5" d="M2 9H22" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                                                                            </svg>
                                                                                            Jan 20, 1989
                                                                                        </li>
                                                                                        <li className="flex items-center gap-2">
                                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                                <path
                                                                                                    opacity="0.5"
                                                                                                    d="M5 8.51464C5 4.9167 8.13401 2 12 2C15.866 2 19 4.9167 19 8.51464C19 12.0844 16.7658 16.2499 13.2801 17.7396C12.4675 18.0868 11.5325 18.0868 10.7199 17.7396C7.23416 16.2499 5 12.0844 5 8.51464Z"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                />
                                                                                                <path
                                                                                                    d="M14 9C14 10.1046 13.1046 11 12 11C10.8954 11 10 10.1046 10 9C10 7.89543 10.8954 7 12 7C13.1046 7 14 7.89543 14 9Z"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                />
                                                                                                <path
                                                                                                    d="M20.9605 15.5C21.6259 16.1025 22 16.7816 22 17.5C22 19.9853 17.5228 22 12 22C6.47715 22 2 19.9853 2 17.5C2 16.7816 2.37412 16.1025 3.03947 15.5"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                    strokeLinecap="round"
                                                                                                />
                                                                                            </svg>
                                                                                            New York, USA
                                                                                        </li>
                                                                                        <li>
                                                                                            <button className="flex items-center gap-2">
                                                                                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                                                                    <path
                                                                                                        opacity="0.5"
                                                                                                        d="M2 12C2 8.22876 2 6.34315 3.17157 5.17157C4.34315 4 6.22876 4 10 4H14C17.7712 4 19.6569 4 20.8284 5.17157C22 6.34315 22 8.22876 22 12C22 15.7712 22 17.6569 20.8284 18.8284C19.6569 20 17.7712 20 14 20H10C6.22876 20 4.34315 20 3.17157 18.8284C2 17.6569 2 15.7712 2 12Z"
                                                                                                        stroke="currentColor"
                                                                                                        strokeWidth="1.5"
                                                                                                    />
                                                                                                    <path
                                                                                                        d="M6 8L8.1589 9.79908C9.99553 11.3296 10.9139 12.0949 12 12.0949C13.0861 12.0949 14.0045 11.3296 15.8411 9.79908L18 8"
                                                                                                        stroke="currentColor"
                                                                                                        strokeWidth="1.5"
                                                                                                        strokeLinecap="round"
                                                                                                    />
                                                                                                </svg>
                                                                                                <span className="text-primary">Jimmy@gmail.com</span>
                                                                                            </button>
                                                                                        </li>
                                                                                        <li className="flex items-center gap-2">
                                                                                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                                                                <path
                                                                                                    d="M5.00659 6.93309C5.04956 5.7996 5.70084 4.77423 6.53785 3.93723C7.9308 2.54428 10.1532 2.73144 11.0376 4.31617L11.6866 5.4791C12.2723 6.52858 12.0372 7.90533 11.1147 8.8278M17.067 18.9934C18.2004 18.9505 19.2258 18.2992 20.0628 17.4622C21.4558 16.0692 21.2686 13.8468 19.6839 12.9624L18.5209 12.3134C17.4715 11.7277 16.0947 11.9628 15.1722 12.8853"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                />
                                                                                                <path
                                                                                                    opacity="0.5"
                                                                                                    d="M5.00655 6.93311C4.93421 8.84124 5.41713 12.0817 8.6677 15.3323C11.9183 18.5829 15.1588 19.0658 17.0669 18.9935M15.1722 12.8853C15.1722 12.8853 14.0532 14.0042 12.0245 11.9755C9.99578 9.94676 11.1147 8.82782 11.1147 8.82782"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                />
                                                                                            </svg>
                                                                                            <span className="whitespace-nowrap" dir="ltr">
                                                                                                +1 (530) 555-12121
                                                                                            </span>
                                                                                        </li>
                                                                                    </ul>
                                                                                    <ul className="mt-7 flex items-center justify-center gap-2">
                                                                                        <li>
                                                                                            <button className="btn btn-info flex items-center justify-center rounded-full w-10 h-10 p-0">
                                                                                                <svg
                                                                                                    xmlns="http://www.w3.org/2000/svg"
                                                                                                    width="24px"
                                                                                                    height="24px"
                                                                                                    viewBox="0 0 24 24"
                                                                                                    fill="none"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                    strokeLinecap="round"
                                                                                                    strokeLinejoin="round"
                                                                                                    className="w-5 h-5"
                                                                                                >
                                                                                                    <path d="M23 3a10.9 10.9 0 0 1-3.14 1.53 4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z"></path>
                                                                                                </svg>
                                                                                            </button>
                                                                                        </li>
                                                                                        <li>
                                                                                            <button className="btn btn-danger flex items-center justify-center rounded-full w-10 h-10 p-0">
                                                                                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-5 h-5">
                                                                                                    <path
                                                                                                        d="M3.33946 16.9997C6.10089 21.7826 12.2168 23.4214 16.9997 20.66C18.9493 19.5344 20.3765 17.8514 21.1962 15.9286C22.3875 13.1341 22.2958 9.83304 20.66 6.99972C19.0242 4.1664 16.2112 2.43642 13.1955 2.07088C11.1204 1.81935 8.94932 2.21386 6.99972 3.33946C2.21679 6.10089 0.578039 12.2168 3.33946 16.9997Z"
                                                                                                        stroke="currentColor"
                                                                                                        strokeWidth="1.5"
                                                                                                    />
                                                                                                    <path
                                                                                                        opacity="0.5"
                                                                                                        d="M16.9497 20.5732C16.9497 20.5732 16.0107 13.9821 14.0004 10.5001C11.99 7.01803 7.05018 3.42681 7.05018 3.42681M7.57711 20.8175C9.05874 16.3477 16.4525 11.3931 21.8635 12.5801M16.4139 3.20898C14.926 7.63004 7.67424 12.5123 2.28857 11.4516"
                                                                                                        stroke="currentColor"
                                                                                                        strokeWidth="1.5"
                                                                                                        strokeLinecap="round"
                                                                                                    />
                                                                                                </svg>
                                                                                            </button>
                                                                                        </li>
                                                                                        <li>
                                                                                            <button className="btn btn-dark flex items-center justify-center rounded-full w-10 h-10 p-0">
                                                                                                <svg
                                                                                                    xmlns="http://www.w3.org/2000/svg"
                                                                                                    width="24px"
                                                                                                    height="24px"
                                                                                                    viewBox="0 0 24 24"
                                                                                                    fill="none"
                                                                                                    stroke="currentColor"
                                                                                                    strokeWidth="1.5"
                                                                                                    strokeLinecap="round"
                                                                                                    strokeLinejoin="round"
                                                                                                    className="w-5 h-5"
                                                                                                >
                                                                                                    <path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path>
                                                                                                </svg>
                                                                                            </button>
                                                                                        </li>
                                                                                    </ul>
                                                                                </div>
                                                                            </div>

                                                                            {/* payment details */}
                                                                            <div className="panel">
                                                                                <div className="flex items-center justify-between mb-5">
                                                                                    <h5 className="font-semibold text-lg dark:text-white-light">Card Details</h5>
                                                                                </div>
                                                                                <div>
                                                                                    <div className="border-b border-[#ebedf2] dark:border-[#1b2e4b]">
                                                                                        <div className="flex items-center justify-between py-2">
                                                                                            <div className="flex-none">
                                                                                                <img src="/assets/images/upi-icon.svg" alt="img" />
                                                                                            </div>
                                                                                            <div className="flex items-center justify-between flex-auto ltr:ml-4 rtl:mr-4">
                                                                                                <h6 className="text-[#515365] font-semibold dark:text-white-dark">
                                                                                                    UPI
                                                                                                    <span className="block text-white-dark dark:text-white-light">Unified Payment Interface</span>
                                                                                                </h6>
                                                                                                <span className="badge bg-success ltr:ml-auto rtl:mr-auto">Primary</span>
                                                                                            </div>
                                                                                        </div>
                                                                                        <div className="flex items-center justify-between py-2">
                                                                                            <div className="flex-none">
                                                                                                <img src="/assets/images/card-americanexpress.svg" alt="img" />
                                                                                            </div>
                                                                                            <div className="flex items-center justify-between flex-auto ltr:ml-4 rtl:mr-4">
                                                                                                <h6 className="text-[#515365] font-semibold dark:text-white-dark">
                                                                                                    American Express
                                                                                                    <span className="block text-white-dark dark:text-white-light">Expires on 12/2025</span>
                                                                                                </h6>
                                                                                            </div>
                                                                                        </div>
                                                                                    </div>
                                                                                    <div className="border-b border-[#ebedf2] dark:border-[#1b2e4b]">
                                                                                        <div className="flex items-center justify-between py-2">
                                                                                            <div className="flex-none">
                                                                                                <img src="/assets/images/card-mastercard.svg" alt="img" />
                                                                                            </div>
                                                                                            <div className="flex items-center justify-between flex-auto ltr:ml-4 rtl:mr-4">
                                                                                                <h6 className="text-[#515365] font-semibold dark:text-white-dark">
                                                                                                    Mastercard
                                                                                                    <span className="block text-white-dark dark:text-white-light">Expires on 03/2025</span>
                                                                                                </h6>
                                                                                            </div>
                                                                                        </div>
                                                                                    </div>
                                                                                    <div>
                                                                                        <div className="flex items-center justify-between py-2">
                                                                                            <div className="flex-none">
                                                                                                <img src="/assets/images/card-visa.svg" alt="img" />
                                                                                            </div>
                                                                                            <div className="flex items-center justify-between flex-auto ltr:ml-4 rtl:mr-4">
                                                                                                <h6 className="text-[#515365] font-semibold dark:text-white-dark">
                                                                                                    Visa
                                                                                                    <span className="block text-white-dark dark:text-white-light">Expires on 10/2025</span>
                                                                                                </h6>
                                                                                            </div>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                            </div>

                                                                            {/* plan details */}
                                                                            <div className="panel">
                                                                                <div className="flex items-center mb-10">
                                                                                    <h5 className="font-semibold text-lg dark:text-white-light">Plan Details</h5>                                                                                    
                                                                                </div>
                                                                                <div className="group">
                                                                                    <ul className="list-inside list-disc text-white-dark font-semibold mb-7 space-y-5">
                                                                                        <li>Full Backup</li>
                                                                                        <li>Unlimited Reports</li>
                                                                                        <li>1 Years Data Storage</li>
                                                                                    </ul>
                                                                                    <div className="flex items-center justify-between mb-4 font-semibold">
                                                                                        <p className="flex items-center rounded-full bg-dark px-2 py-1 text-xs text-white-light font-semibold">
                                                                                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="ltr:mr-1 rtl:ml-1">
                                                                                                <circle opacity="0.5" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="1.5" />
                                                                                                <path d="M12 8V12L14.5 14.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
                                                                                            </svg>
                                                                                            5 Days Left
                                                                                        </p>
                                                                                        <p className="text-info">₹ 99 / month</p>
                                                                                    </div>                                                                                    
                                                                                </div>
                                                                                <div className='flex justify-center mt-12'>
                                                                                    <button className="btn btn-secondary rounded-full">Renew Now</button>
                                                                                </div>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            ) : (
                                                                ''
                                                            )}
                                                            {/*  */}
                                                        </div>
                                                    </Dialog.Panel>
                                                </Transition.Child>
                                            </div>
                                        </div>
                                    </Dialog>
                                </Transition>
                            {/* WIZARD MODAL END */}

                        </div>                                

                        {/* Profile */}
                        <div className="dropdown shrink-0 flex">
                            <Dropdown
                                offset={[0, 8]}
                                placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                btnClassName="relative group block"
                                button={<img className="w-9 h-9 rounded-full object-cover saturate-50 group-hover:saturate-100" src="/assets/images/user-profile.jpeg" alt="userProfile" />}
                            >
                                <ul className="text-dark dark:text-white-dark !py-0 w-[230px] font-semibold dark:text-white-light/90">
                                    <li>
                                        <div className="flex items-center px-4 py-4">
                                            <img className="rounded-md w-10 h-10 object-cover" src="/assets/images/user-profile.jpeg" alt="userProfile" />
                                            <div className="ltr:pl-4 rtl:pr-4">
                                                <h4 className="text-base">
                                                    Recepiton Monk
                                                    <span className="text-xs bg-success-light rounded text-success px-1 ltr:ml-2 rtl:ml-2">Portal</span>
                                                </h4>
                                                <button type="button" className="text-black/60 hover:text-primary dark:text-dark-light/60 dark:hover:text-white">
                                                    portal@email.com
                                                </button>
                                            </div>
                                        </div>
                                    </li>
                                    <li>
                                        <Link to="#" className="dark:hover:text-white">
                                            <svg className="ltr:mr-2 rtl:ml-2" width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                <circle cx="12" cy="6" r="4" stroke="currentColor" strokeWidth="1.5" />
                                                <path
                                                    opacity="0.5"
                                                    d="M20 17.5C20 19.9853 20 22 12 22C4 22 4 19.9853 4 17.5C4 15.0147 7.58172 13 12 13C16.4183 13 20 15.0147 20 17.5Z"
                                                    stroke="currentColor"
                                                    strokeWidth="1.5"
                                                />
                                            </svg>
                                            Profile
                                        </Link>
                                    </li>
                                    
                                    <li className="border-t border-white-light dark:border-white-light/10">
                                        <Link to="#" className="text-danger !py-3">
                                            <svg className="ltr:mr-2 rtl:ml-2 rotate-90" width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                <path
                                                    opacity="0.5"
                                                    d="M17 9.00195C19.175 9.01406 20.3529 9.11051 21.1213 9.8789C22 10.7576 22 12.1718 22 15.0002V16.0002C22 18.8286 22 20.2429 21.1213 21.1215C20.2426 22.0002 18.8284 22.0002 16 22.0002H8C5.17157 22.0002 3.75736 22.0002 2.87868 21.1215C2 20.2429 2 18.8286 2 16.0002L2 15.0002C2 12.1718 2 10.7576 2.87868 9.87889C3.64706 9.11051 4.82497 9.01406 7 9.00195"
                                                    stroke="currentColor"
                                                    strokeWidth="1.5"
                                                    strokeLinecap="round"
                                                />
                                                <path d="M12 15L12 2M12 2L15 5.5M12 2L9 5.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
                                            </svg>
                                            Sign Out
                                        </Link>
                                    </li>
                                </ul>
                            </Dropdown>
                        </div>
                    </div>
                </div>                
            </div>
        </header>

      
    );
};

export default Header;

