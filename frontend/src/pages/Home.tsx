import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../store';
import { Link } from 'react-router-dom';
import ReactApexChart from 'react-apexcharts';
import Dropdown from '../components/Dropdown';
import { setPageTitle } from '../store/themeConfigSlice';

const Home = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Home'));
    });
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;
    const [isShowHomeMenu, setIsShowHomeMenu] = useState(false);
    const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;

    // radialbar graph
    const [loading] = useState(false);

    const [options, setOptions] = useState({
        series: [75],
        chart: {
            // height: 200,
            type: 'radialBar',
            toolbar: {
                show: false,
            }
        },
        plotOptions: {
            radialBar: {
                startAngle: -135,
                endAngle: 225,
                hollow: {
                    margin: 0,
                    size: '80%',
                    background: '#fff',
                    image: undefined,
                    imageOffsetX: 0,
                    imageOffsetY: 0,
                    position: 'front',
                    dropShadow: {
                        enabled: true,
                        top: 3,
                        left: 0,
                        blur: 4,
                        opacity: 0.24
                    }
                },
                track: {
                    background: '#fff',
                    strokeWidth: '100%',
                    margin: 0,
                    dropShadow: {
                        enabled: true,
                        top: -3,
                        left: 0,
                        blur: 4,
                        opacity: 0.35
                    }
                },

                dataLabels: {
                    show: true,
                    name: {
                        offsetY: -10,
                        show: true,
                        color: '#444', 
                        fontSize: '16px',
                    },
                    value: {
                        formatter: function (val) {
                            return parseInt(val);
                        },
                        color: '#000',
                        fontSize: '22px',
                        show: true,
                    }
                }
            }
        },
        fill: {
            type: 'solid',
            colors: ['#1937cc'],
        },
        stroke: {
            lineCap: 'round'
        },
        labels: ['Days Left'],
    }
    );

        useEffect(() => {
            // You can set options here if you want to dynamically update them.
        }, []);
    
    // uniqueVisitorSeriesOptions
    const uniqueVisitorSeries: any = {
        series: [
            {
                name: 'Answered',
                data: [58, 44, 55, 57, 56, 61, 58, 63, 60, 66, 56, 63],
            },
            {
                name: 'Missed',
                data: [91, 76, 85, 101, 98, 87, 105, 91, 114, 94, 66, 70],
            },
        ],
        options: {
            chart: {
                height: 360,
                type: 'bar',
                fontFamily: 'Nunito, sans-serif',
                toolbar: {
                    show: false,
                },
            },
            dataLabels: {
                enabled: false,
            },
            stroke: {
                width: 2,
                colors: ['transparent'],
            },
            colors: ['#5c1ac3', '#ffbb44'],
            dropShadow: {
                enabled: true,
                blur: 3,
                color: '#515365',
                opacity: 0.4,
            },
            plotOptions: {
                bar: {
                    horizontal: false,
                    columnWidth: '55%',
                    borderRadius: 8,
                    borderRadiusApplication: 'end',
                },
            },
            legend: {
                position: 'bottom',
                horizontalAlign: 'center',
                fontSize: '14px',
                itemMargin: {
                    horizontal: 8,
                    vertical: 8,
                },
            },
            grid: {
                borderColor: isDark ? '#191e3a' : '#e0e6ed',
                padding: {
                    left: 20,
                    right: 20,
                },
                xaxis: {
                    lines: {
                        show: false,
                    },
                },
            },
            xaxis: {
                categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
                axisBorder: {
                    show: true,
                    color: isDark ? '#3b3f5c' : '#e0e6ed',
                },
            },
            yaxis: {
                tickAmount: 6,
                opposite: isRtl ? true : false,
                labels: {
                    offsetX: isRtl ? -10 : 0,
                },
            },
            fill: {
                type: 'gradient',
                gradient: {
                    shade: isDark ? 'dark' : 'light',
                    type: 'vertical',
                    shadeIntensity: 0.3,
                    inverseColors: false,
                    opacityFrom: 1,
                    opacityTo: 0.8,
                    stops: [0, 100],
                },
            },
            tooltip: {
                marker: {
                    show: true,
                },
            },
        },
    };

    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="/" className="text-primary hover:underline">
                        Dashboard
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Home</span>
                </li>
            </ul>

            <div className="pt-5">
                <div className="grid lg:grid-cols-3 gap-6 mb-6">
                    <div className="grid col-span-2 gap-6">
                        <div className="panel overflow-hidden">
                            <div className="flex items-center justify-between">
                                <div>
                                    <div className="">Welcome Back!</div>
                                    <div className="text-2xl font-bold">My Dashboard</div>
                                </div>
                                <div className="flex">
                                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-4 h-4 me-2">
                                        <path
                                            d="M2 12C2 8.22876 2 6.34315 3.17157 5.17157C4.34315 4 6.22876 4 10 4H14C17.7712 4 19.6569 4 20.8284 5.17157C22 6.34315 22 8.22876 22 12V14C22 17.7712 22 19.6569 20.8284 20.8284C19.6569 22 17.7712 22 14 22H10C6.22876 22 4.34315 22 3.17157 20.8284C2 19.6569 2 17.7712 2 14V12Z"
                                            stroke="currentColor"
                                            strokeWidth="1.5"
                                        />
                                        <path opacity="0.5" d="M7 4V2.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M17 4V2.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M2 9H22" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                    </svg>
                                    <div className="text-xs font-bold">                                        
                                        12 April, 2023
                                    </div>
                                </div>
                                <div className="flex">
                                    <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="16" height="16" viewBox="0 0 48 48" className="w-4 h-4 me-2">
                                        <path d="M 24 4 C 12.972066 4 4 12.972074 4 24 C 4 35.027926 12.972066 44 24 44 C 35.027934 44 44 35.027926 44 24 C 44 12.972074 35.027934 4 24 4 z M 24 7 C 33.406615 7 41 14.593391 41 24 C 41 33.406609 33.406615 41 24 41 C 14.593385 41 7 33.406609 7 24 C 7 14.593391 14.593385 7 24 7 z M 23.476562 11.978516 A 1.50015 1.50015 0 0 0 22 13.5 L 22 25.5 A 1.50015 1.50015 0 0 0 23.5 27 L 31.5 27 A 1.50015 1.50015 0 1 0 31.5 24 L 25 24 L 25 13.5 A 1.50015 1.50015 0 0 0 23.476562 11.978516 z"></path>
                                    </svg>
                                    <div className="text-xs font-bold">                                        
                                        15:25:56
                                    </div>
                                </div>
                            </div>
                            <div className="relative mt-10">
                                <div className="grid grid-cols-3 gap-6 p-2 text-center">                                    
                                    <div className="flex flex-col justify-center items-center">
                                        <img src="/assets/images/profile-34.jpeg" alt="img" className="w-24 h-24 rounded-full object-cover  mb-5" />
                                        <p className="font-semibold text-primary text-2xl">John Miller</p>
                                    </div>

                                    <ul className="flex flex-col max-w-[160px] m-auto space-y-4 font-semibold text-gray-600">
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
                                            John@gmail.com
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
                                        <span className="whitespace-nowrap text-primary font-bold" dir="ltr">
                                            +91 98562 52145
                                        </span>
                                    </li>
                                    </ul>

                                    <div>
                                        <ReactApexChart options={options} series={options.series} type="radialBar" height={200} /> 

                                        {/* <span className="font-semibold text-dark text-xl" dir="ltr">
                                            Basic Plan
                                        </span> */}
                                    </div>
                                </div>                               
                            </div>
                        </div>
                    </div>

                    <div className="relative">
                    <div className="panel absolute inset-0 bg-gradient-to-r from-violet-500 to-violet-400"></div>
                        <div
                            className="absolute inset-0"
                            className={{ backgroundImage: `url('/assets/images/bg/circle-vector.png')`, backgroundRepeat: 'no-repeat', backgroundSize: 'cover' }}>
                            <div className="flex items-center justify-center p-8">
                                <svg className="svg-main-icon" id="Capa_1" enable-background="new 0 0 512 512" viewBox="0 0 512 512" width="50" height="50" xmlns="http://www.w3.org/2000/svg">
                                    <path d="m482 245.242v-60.363c0-29.656-23.597-53.891-53-54.949v-37.051c0-19.299-15.701-35-35-35h-96.358l-12.443-34.587c-3.173-8.82-9.595-15.868-18.083-19.845-8.488-3.978-18.014-4.402-26.821-1.196l-174.855 63.641c-8.798 3.202-15.817 9.641-19.765 18.131s-4.349 18.007-1.128 26.799l7.025 19.175c-28.735 1.777-51.572 25.707-51.572 54.882v272c0 30.327 24.673 55 55 55h372c30.327 0 55-24.673 55-55v-62.363c16.938-2.434 30-17.036 30-34.637v-80c0-17.601-13.062-32.203-30-34.637zm0 114.637c0 2.757-2.243 5-5 5h-80c-24.813 0-45-20.187-45-45s20.187-45 45-45h80c2.757 0 5 2.243 5 5zm-409.284-259.377c-.621-1.695-.166-3.126.161-3.829.327-.702 1.128-1.973 2.824-2.59l174.854-63.641c1.698-.617 3.129-.158 3.832.171s1.972 1.135 2.583 2.835l8.79 24.432h-6.76c-19.299 0-35 15.701-35 35v37h-140.521zm326.284-7.623v37h-145v-37c0-2.757 2.243-5 5-5h135c2.757 0 5 2.243 5 5zm28 389h-372c-13.785 0-25-11.215-25-25v-272c0-13.785 11.215-25 25-25h372c13.785 0 25 11.215 25 25v60h-55c-41.355 0-75 33.645-75 75s33.645 75 75 75h55v62c0 13.785-11.215 25-25 25z" fill="#fff"></path>
                                    <circle cx="397" cy="319.879" r="15" fill="#fff"></circle>
                                </svg>
                                <div className="font-bold text-3xl text-white ps-6 pt-2">Wallet Balance</div>                          
                            </div>
                            <div className="">
                                <div className="flex justify-center align-center whitespace-nowrap py-3">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="#fff" className="mt-2"><path d="M15.571 3h3.866l1.563-3h-16.438l-1.562 3h4.963c2.257 0 3.759.839 4.589 2h-7.99l-1.562 3h10.279c-.266 2.071-1.975 4-4.942 4h-4.337v3c2.321 0 1.584-.553 7.324 9h4.676l-5.963-9c2.505-.396 6.496-2.415 6.92-7h2.48l1.563-3h-4.345c-.238-.707-.602-1.383-1.084-2z"/></svg>
                                    <span className="text-3xl text-white font-bold">41,741.42</span>                                
                                </div>                            
                            </div>
                            <div className="flex items-end justify-evenly">
                                <div className="pt-12">
                                    <button type="button" className="flex shadow-[0_0_2px_0_#bfc9d4] bg-[#EBF1F6] rounded-full px-6 py-2 text-white-light place-content-center ltr:mr-2 rtl:ml-2">
                                        <span className='font-semibold text-lg text-gray-900'>Purchase Plan</span>
                                    </button>
                                </div>
                                <div className="pt-12">
                                    <button type="button" className="btn btn-primary rounded shadow-[0_0_1px_0_#bfc9d4] rounded-full ps-6 text-white text-base hover:bg-[#4361ee]">
                                        Recharge Wallet
                                        <span className='bg-[#EBF1F6] rounded-xl ms-3'>
                                            <svg className="w-6 h-6" viewBox="0 0 24 24" stroke="#333333" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
                                                <line x1="12" y1="5" x2="12" y2="19"></line>
                                                <line x1="5" y1="12" x2="19" y2="12"></line>
                                            </svg>
                                        </span>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
                <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
                    <div className="panel h-full sm:col-span-2 lg:col-span-1 shadow-md">                       
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Answered Calls</h5>                            
                        </div>
                        
                        <div className="grid sm:grid-cols-2 gap-8 text-sm text-[#515365] font-bold">
                            <div>
                                <div className='bg-[#04DC8B] flex align-center rounded-full justify-center w-28 h-28 '>
                                    <div className="grid place-content-center">                                
                                    <svg width="50" height="50" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path d="M14.5 6.5C15.2372 6.64382 15.9689 6.96892 16.5 7.5C17.0311 8.03108 17.3562 8.76284 17.5 9.5M15 3C16.5315 3.17014 17.9097 3.91107 19 5C20.0903 6.08893 20.8279 7.46869 21 9M20.9995 16.4767V19.1864C21.0037 20.2223 20.0723 21.0873 19.0265 20.9929C10.0001 21 3.00006 13.935 3.00713 4.96919C2.91294 3.92895 3.77364 3.00106 4.80817 3.00009H7.52331C7.96253 2.99577 8.38835 3.151 8.72138 3.43684C9.66819 4.24949 10.2772 7.00777 10.0429 8.10428C9.85994 8.96036 8.99696 9.55929 8.41026 10.1448C9.69864 12.4062 11.5747 14.2785 13.8405 15.5644C14.4272 14.9788 15.0274 14.1176 15.8851 13.935C16.9855 13.7008 19.7615 14.3106 20.5709 15.264C20.858 15.6021 21.0105 16.0337 20.9995 16.4767Z" stroke="#fff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                                    </svg>
                                    </div>
                                </div>
                            </div>

                            <div className='flex items-center'>
                                <div>
                                    <div>Total Received</div>
                                    <div className="text-[#04DC8B] text-3xl">7,929</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className="panel h-full sm:col-span-2 lg:col-span-1 shadow-md">
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Missed Calls</h5>
                        </div>

                        <div className="grid sm:grid-cols-2 gap-8 text-sm text-[#515365] font-bold">
                            <div>
                                <div className='bg-[#dc3545] flex align-center rounded-full justify-center w-28 h-28 '>
                                    <div className="grid place-content-center">                                
                                    <svg width="50" height="50" viewBox="0 0 64 64" xmlns="http://www.w3.org/2000/svg" stroke-width="3" stroke="#ffffff" fill="none">
                                        <path d="M11.11,8.4a2.62,2.62,0,0,0-2.53,2.78c.35,6,2,20.64,9.9,29.77,9.46,11,21.78,14.79,34.42,14.23a2.68,2.68,0,0,0,2.52-2.65V42.92a4,4,0,0,0-3.09-3.86L46,37.66a4,4,0,0,0-4.16,1.69l-1.4,2.12a1,1,0,0,1-1.22.37C36,40.45,23.17,34.45,21.76,24.33a1,1,0,0,1,.48-1l2.54-1.55a4,4,0,0,0,1.81-4.21L25.2,11.13a4,4,0,0,0-4-3.12Z"/><line x1="39.32" y1="10.89" x2="53.65" y2="25.22"/><line x1="39.32" y1="25.22" x2="53.65" y2="10.89"/>
                                    </svg>
                                    </div>
                                </div>
                            </div>

                            <div className='flex items-center'>
                                <div>
                                    <div>Total Missed</div>
                                    <div className="text-[#dc3545] text-3xl">1,419</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className="panel h-full sm:col-span-2 lg:col-span-1 shadow-md">
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Device Registered</h5>
                        </div>

                        <div className="grid sm:grid-cols-2 gap-8 text-sm text-[#515365] font-bold">
                            <div>
                                <div className='bg-[#0dcaf0] flex align-center rounded-full justify-center w-28 h-28 '>
                                    <div className="grid place-content-center">                                
                                    <svg height="50" width="50" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
                                        viewBox="0 0 295.239 295.239" xml:space="preserve">
                                    <g>
                                        <g>
                                            <path fill="#ffffff" d="M147.62,166.667c-57.767,0-104.762,46.995-104.762,104.762v23.81h209.524v-23.81
                                                C252.382,213.662,205.387,166.667,147.62,166.667z M189.782,186.162c-13.552,29.305-33.367,71.714-41.9,89.029l-42.4-89.043
                                                c12.724-6.314,26.995-9.957,42.138-9.957C162.772,176.19,177.053,179.838,189.782,186.162z M196.901,193.467l11.948,51.771
                                                l-49.533,28.305C165.796,260.395,176.958,236.6,196.901,193.467z M136.786,274.038l-50.396-28.795l11.976-51.886L136.786,274.038z
                                                M52.382,271.429c0-29.786,13.762-56.395,35.238-73.871l-12.1,52.438l62.505,35.719H52.382V271.429z M242.858,285.714h-85.643
                                                l62.509-35.719l-12.105-52.438c21.476,17.476,35.238,44.086,35.238,73.871v14.286H242.858z"/>
                                            <path fill="#ffffff" d="M207.001,138.095h12.048c7.876,0,14.286-6.41,14.286-14.286v-14.286h-9.524v14.286
                                                c0,2.624-2.133,4.762-4.762,4.762h-5.548c6.514-11.224,10.31-24.21,10.31-38.095c0-42.01-34.181-76.19-76.19-76.19
                                                s-76.19,34.181-76.19,76.19s34.181,76.19,76.19,76.19C171.615,166.667,193.025,155.49,207.001,138.095z M80.953,90.476
                                                c0-36.762,29.905-66.667,66.667-66.667s66.667,29.905,66.667,66.667c0,14.162-4.471,27.286-12.033,38.095h-31.7
                                                c-1.971-5.529-7.21-9.524-13.41-9.524h-19.048c-7.876,0-14.286,6.41-14.286,14.286s6.41,14.286,14.286,14.286h19.048
                                                c6.2,0,11.438-3.995,13.41-9.524h23.638c-12.029,11.762-28.457,19.048-46.567,19.048
                                                C110.858,157.143,80.953,127.238,80.953,90.476z M161.906,133.333c0,2.624-2.133,4.762-4.762,4.762h-19.048
                                                c-2.629,0-4.762-2.138-4.762-4.762s2.133-4.762,4.762-4.762h19.048C159.772,128.571,161.906,130.71,161.906,133.333z"/>
                                            <path fill="#ffffff" d="M147.62,9.524c34.414,0,65.138,21.814,76.457,54.281l8.995-3.133
                                                C220.415,24.381,186.077,0,147.62,0c-38.452,0-72.79,24.381-85.448,60.667l8.995,3.138C82.487,31.338,113.211,9.524,147.62,9.524z
                                                "/>
                                            <path fill="#333333" d="M71.43,123.81h-4.762c-18.376,0-33.333-14.957-33.333-33.333s14.957-33.333,33.333-33.333h4.762
                                                V123.81z"/>
                                            <path fill="#333333" d="M228.572,123.81h-4.762V57.143h4.762c18.376,0,33.333,14.957,33.333,33.333
                                                S246.949,123.81,228.572,123.81z"/>
                                        </g>
                                    </g>
                                    </svg>
                                    </div>
                                </div>
                            </div>

                            <div className='flex items-center'>
                                <div>
                                    <div>Total IVR</div>
                                    <div className="text-[#0dcaf0] text-3xl">5,233</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className="panel h-full sm:col-span-2 lg:col-span-1 shadow-md">
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Extensions</h5>
                        </div>

                        <div className="grid sm:grid-cols-2 gap-8 text-sm text-[#515365] font-bold">
                            <div>
                                <div className='bg-[#FD8F01] flex align-center rounded-full justify-center w-28 h-28 '>
                                    <div className="grid place-content-center">                                
                                    {/* <svg fill="#ffffff" width="50" height="50" viewBox="0 0 256 256" id="Layer_1" version="1.1" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">

                                        <g>

                                        <path d="M184.4,48.5L184.4,48.5c1.9,9.7,10.4,16.8,20.3,16.8c9.9,0,18.4-7.1,20.3-16.8l2.5-12.7c0.8-6.2-1-12.4-5.1-17.1   s-10-7.4-16.2-7.4h-2.9c-6.2,0-12.1,2.7-16.2,7.4s-6,10.9-5.1,17.4L184.4,48.5z M194.6,25.3c2.2-2.5,5.4-4,8.8-4h2.9   c3.4,0,6.5,1.4,8.8,4c2.2,2.5,3.2,5.9,2.8,8.9l-2.4,12.4c-1,5.1-5.4,8.7-10.6,8.7s-9.6-3.7-10.6-8.7l-2.4-12.1   C191.4,31.2,192.4,27.8,194.6,25.3z"/>

                                        <path d="M251.3,114.4l-5.7-27c-0.7-3.3-3.1-6-6.2-7.1l-19.7-7l-14.8,14.8l-14.8-14.8l-19.7,7c-3.1,1.1-5.5,3.8-6.2,7.1l-5.7,27   c-0.6,2.7,1.1,5.3,3.8,5.9c2.7,0.6,5.3-1.1,5.9-3.8l5.6-26.8l13.8-4.9l17.4,17.4l17.4-17.4l13.6,4.7l5.7,27   c0.5,2.3,2.6,3.9,4.9,3.9c0.3,0,0.7,0,1-0.1C250.1,119.8,251.8,117.1,251.3,114.4z"/>

                                        <path d="M30.9,48.5L30.9,48.5c1.9,9.7,10.4,16.8,20.3,16.8c9.9,0,18.4-7.1,20.3-16.8L74,35.8c0.8-6.2-1-12.4-5.1-17.1   s-10-7.4-16.2-7.4h-2.9c-6.2,0-12.1,2.7-16.2,7.4s-6,10.9-5.1,17.4L30.9,48.5z M41.1,25.3c2.2-2.5,5.4-4,8.8-4h2.9   c3.4,0,6.5,1.4,8.8,4c2.2,2.5,3.2,5.9,2.8,8.9l-2.4,12.4c-1,5.1-5.4,8.7-10.6,8.7s-9.6-3.7-10.6-8.7l-2.4-12.1   C37.8,31.2,38.9,27.8,41.1,25.3z"/>

                                        <path d="M68.7,84.8l13.6,4.7l5.7,27c0.5,2.3,2.6,3.9,4.9,3.9c0.3,0,0.7,0,1-0.1c2.7-0.6,4.4-3.2,3.8-5.9l-5.7-27   c-0.7-3.3-3.1-6-6.2-7.1l-19.7-7L51.2,88.2L36.4,73.3l-19.7,7c-3.1,1.1-5.5,3.8-6.2,7.1l-5.7,27c-0.6,2.7,1.1,5.3,3.8,5.9   c2.7,0.6,5.3-1.1,5.9-3.8L20,89.7l13.8-4.9l17.4,17.4L68.7,84.8z"/>

                                        <path d="M128,189.5c9.9,0,18.4-7.1,20.3-16.8l2.5-12.7c0.8-6.2-1-12.4-5.1-17.1c-4.1-4.7-10-7.4-16.2-7.4h-2.9   c-6.2,0-12.1,2.7-16.2,7.4c-4.1,4.7-6,10.9-5.1,17.4l2.4,12.4v0C109.6,182.4,118.1,189.5,128,189.5z M117.8,149.5   c2.2-2.5,5.4-4,8.8-4h2.9c3.4,0,6.5,1.4,8.8,4c2.2,2.5,3.2,5.9,2.8,8.9l-2.4,12.4c-1,5.1-5.4,8.7-10.6,8.7s-9.6-3.7-10.6-8.7   l-2.4-12.1C114.6,155.4,115.6,152,117.8,149.5z"/>

                                        <path d="M168.8,211.7c-0.7-3.3-3.1-6-6.2-7.1l-19.7-7L128,212.4l-14.8-14.8l-19.7,7c-3.1,1.1-5.5,3.8-6.2,7.1l-5.7,27   c-0.6,2.7,1.1,5.3,3.8,5.9c2.7,0.6,5.3-1.1,5.9-3.8l5.5-26.8l13.8-4.9l17.4,17.4l17.4-17.4l13.6,4.7l5.7,27   c0.5,2.3,2.6,3.9,4.9,3.9c0.3,0,0.7,0,1-0.1c2.7-0.6,4.4-3.2,3.8-5.9L168.8,211.7z"/>

                                        <path d="M213.4,133.5c-2.8-0.3-5.2,1.7-5.4,4.5c-2,20.7-11.9,39.5-27.7,53c-2.1,1.8-2.3,4.9-0.6,7c1,1.2,2.4,1.7,3.8,1.7   c1.1,0,2.3-0.4,3.2-1.2c17.8-15.2,28.9-36.3,31.2-59.6C218.1,136.2,216.1,133.8,213.4,133.5z"/>

                                        <path d="M48.1,138c-0.3-2.7-2.7-4.7-5.4-4.5c-2.7,0.3-4.7,2.7-4.5,5.4c2.3,23.3,13.4,44.5,31.2,59.6c0.9,0.8,2.1,1.2,3.2,1.2   c1.4,0,2.8-0.6,3.8-1.8c1.8-2.1,1.5-5.2-0.6-7C60,177.5,50.1,158.7,48.1,138z"/>

                                        <path d="M101.7,54.1c16.7-5.9,35.9-5.9,52.7,0c0.5,0.2,1.1,0.3,1.7,0.3c2,0,4-1.3,4.7-3.3c0.9-2.6-0.4-5.4-3-6.3   c-18.9-6.7-40.4-6.7-59.3,0c-2.6,0.9-3.9,3.8-3,6.3C96.2,53.7,99.1,55,101.7,54.1z"/>

                                        </g>

                                        </svg> */}

                                        <svg width="50" height="50" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <g opacity="0.5">
                                            <path d="M7.20468 7.56232C7.51523 7.28822 7.54478 6.81427 7.27069 6.50371C6.99659 6.19316 6.52264 6.1636 6.21208 6.4377C4.39676 8.03992 3.25 10.3865 3.25 13C3.25 13.4142 3.58579 13.75 4 13.75C4.41421 13.75 4.75 13.4142 4.75 13C4.75 10.8347 5.69828 8.89188 7.20468 7.56232Z" fill="#1C274C"/>
                                            <path d="M17.7879 6.4377C17.4774 6.1636 17.0034 6.19316 16.7293 6.50371C16.4552 6.81427 16.4848 7.28822 16.7953 7.56232C18.3017 8.89188 19.25 10.8347 19.25 13C19.25 13.4142 19.5858 13.75 20 13.75C20.4142 13.75 20.75 13.4142 20.75 13C20.75 10.3865 19.6032 8.03992 17.7879 6.4377Z" fill="#1C274C"/>
                                            <path d="M10.1869 20.0217C9.7858 19.9184 9.37692 20.1599 9.27367 20.5611C9.17043 20.9622 9.41192 21.3711 9.81306 21.4743C10.5129 21.6544 11.2458 21.75 12 21.75C12.7542 21.75 13.4871 21.6544 14.1869 21.4743C14.5881 21.3711 14.8296 20.9622 14.7263 20.5611C14.6231 20.1599 14.2142 19.9184 13.8131 20.0217C13.2344 20.1706 12.627 20.25 12 20.25C11.373 20.25 10.7656 20.1706 10.1869 20.0217Z" fill="#1C274C"/>
                                            </g>
                                            <path d="M9 6C9 7.65685 10.3431 9 12 9C13.6569 9 15 7.65685 15 6C15 4.34315 13.6569 3 12 3C10.3431 3 9 4.34315 9 6Z" fill="#1C274C"/>
                                            <path d="M2.5 18C2.5 19.6569 3.84315 21 5.5 21C7.15685 21 8.5 19.6569 8.5 18C8.5 16.3431 7.15685 15 5.5 15C3.84315 15 2.5 16.3431 2.5 18Z" fill="#1C274C"/>
                                            <path d="M18.5 21C16.8431 21 15.5 19.6569 15.5 18C15.5 16.3431 16.8431 15 18.5 15C20.1569 15 21.5 16.3431 21.5 18C21.5 19.6569 20.1569 21 18.5 21Z" fill="#1C274C"/>
                                        </svg>
                                    </div>
                                </div>
                            </div>

                            <div className='flex items-center'>
                                <div>
                                    <div>Total Members</div>
                                    <div className="text-[#FD8F01] text-3xl">10,419</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                        
            <div className="grid col-span-12">
                <div className="panel h-full">
                    <div className="grid grid-cols-12">
                        <div className="grid col-span-9">                            
                            <div className="p-2">
                                <ReactApexChart options={uniqueVisitorSeries.options} series={uniqueVisitorSeries.series} type="bar" height={360} />
                            </div>
                        </div>
                        <div className="grid col-span-3">
                            <div className="">
                                <div className="flex justify-between">
                                    <select className="image-select default-select dashboard-select primary-light style-1 hidden" aria-label="Default">
                                        <option selected>This Month</option>
                                        <option value="1">This Week</option>
                                        <option value="2">This Day</option>
                                    </select>
                                    <div className="nice-select image-select default-select dashboard-select primary-light style-1" tabindex="0">
                                        <span className="current">This Month</span>
                                        <ul className="list">
                                            <li data-value="This Month" className="option selected">This Month</li>
                                            <li data-value="1" className="option">This Week</li>
                                            <li data-value="2" className="option">This Day</li>
                                        </ul>
                                    </div>                                    
                                </div>
                            </div>
                            <div className="card expense mb-3">
                                <div className="card-body p-3">
                                    <div className="students1 flex items-center justify-between">
                                        <div className="content">
                                            <span>Answered</span>
                                            <h2>7,929</h2>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="card expense mb-3">
                                <div className="card-body p-3">
                                    <div className="students1 flex items-center justify-between">
                                        <div className="content">
                                            <span>Missed</span>
                                            <h2>1,419</h2>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    );
};

export default Home;