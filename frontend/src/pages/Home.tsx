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
    // totalVisitOptions
    const totalVisit: any = {
        series: [{ data: [21, 9, 36, 12, 44, 25, 59, 41, 66, 25] }],
        options: {
            chart: {
                height: 58,
                type: 'line',
                fontFamily: 'Nunito, sans-serif',
                sparkline: {
                    enabled: true,
                },
                dropShadow: {
                    enabled: true,
                    blur: 3,
                    color: '#009688',
                    opacity: 0.4,
                },
            },
            stroke: {
                curve: 'smooth',
                width: 2,
            },
            colors: ['#009688'],
            grid: {
                padding: {
                    top: 5,
                    bottom: 5,
                    left: 5,
                    right: 5,
                },
            },
            tooltip: {
                x: {
                    show: false,
                },
                y: {
                    title: {
                        formatter: () => {
                            return '';
                        },
                    },
                },
            },
        },
    };
    // paidVisitOptions
    const paidVisit: any = {
        series: [{ data: [22, 19, 30, 47, 32, 44, 34, 55, 41, 69] }],
        options: {
            chart: {
                height: 58,
                type: 'line',
                fontFamily: 'Nunito, sans-serif',
                sparkline: {
                    enabled: true,
                },
                dropShadow: {
                    enabled: true,
                    blur: 3,
                    color: '#e2a03f',
                    opacity: 0.4,
                },
            },
            stroke: {
                curve: 'smooth',
                width: 2,
            },
            colors: ['#e2a03f'],
            grid: {
                padding: {
                    top: 5,
                    bottom: 5,
                    left: 5,
                    right: 5,
                },
            },
            tooltip: {
                x: {
                    show: false,
                },
                y: {
                    title: {
                        formatter: () => {
                            return '';
                        },
                    },
                },
            },
        },
    };

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
                                <div className="grid grid-cols-3 gap-6 p-2">                                    
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
                                    </div>
                                </div>                               
                            </div>
                        </div>
                    </div>

                    <div className="panel bg-gradient-to-r from-violet-500 to-violet-400">                        
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
                        <div className="flex items-center justify-evenly">
                                <div className="py-8">
                                    <button type="button" className="shadow-[0_0_2px_0_#bfc9d4] rounded p-1 text-white-light hover:bg-[#1937cc] place-content-center ltr:mr-2 rtl:ml-2">
                                        <svg className="w-8 h-8" viewBox="0 0 24 24" stroke="currentColor" strokeWidth="1.5" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                            <line x1="12" y1="5" x2="12" y2="19"></line>
                                            <line x1="5" y1="12" x2="19" y2="12"></line>
                                        </svg>
                                    </button>
                                </div>
                                <button type="button" className="btn btn-primary rounded shadow-[0_0_1px_0_#bfc9d4] rounded px-6 text-white text-base hover:bg-[#4361ee] z-10">
                                    Recharge
                                </button>
                            </div>
                    </div>
                </div>

                <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
                    <div className="panel h-full sm:col-span-2 lg:col-span-1">
                        {/* statistics */}
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Statistics</h5>
                            <div className="dropdown">
                                <Dropdown
                                    offset={[0, 5]}
                                    placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                    btnClassName="hover:text-primary"
                                    button={
                                        <svg className="w-5 h-5 text-black/70 dark:text-white/70 hover:!text-primary" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <circle cx="5" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle opacity="0.5" cx="12" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle cx="19" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                        </svg>
                                    }
                                >
                                    <ul>
                                        <li>
                                            <button type="button">This Week</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Week</button>
                                        </li>
                                        <li>
                                            <button type="button">This Month</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Month</button>
                                        </li>
                                    </ul>
                                </Dropdown>
                            </div>
                        </div>
                        <div className="grid sm:grid-cols-2 gap-8 text-sm text-[#515365] font-bold">
                            <div>
                                <div>
                                    <div>Total Visits</div>
                                    <div className="text-[#f8538d] text-lg">423,964</div>
                                </div>

                                <ReactApexChart series={totalVisit.series} options={totalVisit.options} type="line" height={58} />
                            </div>

                            <div>
                                <div>
                                    <div>Paid Visits</div>
                                    <div className="text-[#f8538d] text-lg">7,929</div>
                                </div>

                                <ReactApexChart series={paidVisit.series} options={paidVisit.options} type="line" height={58} />
                            </div>
                        </div>
                    </div>

                    <div className="panel h-full">
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Expenses</h5>

                            <div className="dropdown">
                                <Dropdown
                                    offset={[0, 5]}
                                    placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                    btnClassName="hover:text-primary"
                                    button={
                                        <svg className="w-5 h-5 text-black/70 dark:text-white/70 hover:!text-primary" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <circle cx="5" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle opacity="0.5" cx="12" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle cx="19" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                        </svg>
                                    }
                                >
                                    <ul>
                                        <li>
                                            <button type="button">This Week</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Week</button>
                                        </li>
                                        <li>
                                            <button type="button">This Month</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Month</button>
                                        </li>
                                    </ul>
                                </Dropdown>
                            </div>
                        </div>
                        <div className=" text-[#e95f2b] text-3xl font-bold my-10">
                            <span>$ 45,141 </span>
                            <span className="text-black text-sm dark:text-white-light ltr:mr-2 rtl:ml-2">this week</span>
                            <svg className="text-success inline" width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path
                                    opacity="0.5"
                                    d="M22 7L14.6203 14.3347C13.6227 15.3263 13.1238 15.822 12.5051 15.822C11.8864 15.8219 11.3876 15.326 10.3902 14.3342L10.1509 14.0962C9.15254 13.1035 8.65338 12.6071 8.03422 12.6074C7.41506 12.6076 6.91626 13.1043 5.91867 14.0977L2 18"
                                    stroke="currentColor"
                                    strokeWidth="1.5"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                />
                                <path d="M22.0001 12.5458V7H16.418" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
                            </svg>
                        </div>
                        <div className="flex items-center justify-between">
                            <div className="w-full rounded-full h-5 p-1 bg-dark-light overflow-hidden shadow-3xl dark:shadow-none dark:bg-dark-light/10">
                                <div
                                    className="bg-gradient-to-r from-[#4361ee] to-[#805dca] w-full h-full rounded-full relative before:absolute before:inset-y-0 ltr:before:right-0.5 rtl:before:left-0.5 before:bg-white before:w-2 before:h-2 before:rounded-full before:m-auto"
                                    style={{ width: '65%' }}
                                ></div>
                            </div>
                            <span className="ltr:ml-5 rtl:mr-5 dark:text-white-light">57%</span>
                        </div>
                    </div>

                    <div className="panel h-full sm:col-span-2 lg:col-span-1">
                        {/* statistics */}
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Statistics</h5>
                            <div className="dropdown">
                                <Dropdown
                                    offset={[0, 5]}
                                    placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                    btnClassName="hover:text-primary"
                                    button={
                                        <svg className="w-5 h-5 text-black/70 dark:text-white/70 hover:!text-primary" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <circle cx="5" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle opacity="0.5" cx="12" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle cx="19" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                        </svg>
                                    }
                                >
                                    <ul>
                                        <li>
                                            <button type="button">This Week</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Week</button>
                                        </li>
                                        <li>
                                            <button type="button">This Month</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Month</button>
                                        </li>
                                    </ul>
                                </Dropdown>
                            </div>
                        </div>
                        <div className="grid sm:grid-cols-2 gap-8 text-sm text-[#515365] font-bold">
                            <div>
                                <div>
                                    <div>Total Visits</div>
                                    <div className="text-[#f8538d] text-lg">423,964</div>
                                </div>

                                <ReactApexChart series={totalVisit.series} options={totalVisit.options} type="line" height={58} />
                            </div>

                            <div>
                                <div>
                                    <div>Paid Visits</div>
                                    <div className="text-[#f8538d] text-lg">7,929</div>
                                </div>

                                <ReactApexChart series={paidVisit.series} options={paidVisit.options} type="line" height={58} />
                            </div>
                        </div>
                    </div>

                    <div className="panel h-full">
                        <div className="flex justify-between dark:text-white-light mb-5">
                            <h5 className="font-semibold text-lg ">Expenses</h5>

                            <div className="dropdown">
                                <Dropdown
                                    offset={[0, 5]}
                                    placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                    btnClassName="hover:text-primary"
                                    button={
                                        <svg className="w-5 h-5 text-black/70 dark:text-white/70 hover:!text-primary" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <circle cx="5" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle opacity="0.5" cx="12" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                            <circle cx="19" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                        </svg>
                                    }
                                >
                                    <ul>
                                        <li>
                                            <button type="button">This Week</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Week</button>
                                        </li>
                                        <li>
                                            <button type="button">This Month</button>
                                        </li>
                                        <li>
                                            <button type="button">Last Month</button>
                                        </li>
                                    </ul>
                                </Dropdown>
                            </div>
                        </div>
                        <div className=" text-[#e95f2b] text-3xl font-bold my-10">
                            <span>$ 45,141 </span>
                            <span className="text-black text-sm dark:text-white-light ltr:mr-2 rtl:ml-2">this week</span>
                            <svg className="text-success inline" width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path
                                    opacity="0.5"
                                    d="M22 7L14.6203 14.3347C13.6227 15.3263 13.1238 15.822 12.5051 15.822C11.8864 15.8219 11.3876 15.326 10.3902 14.3342L10.1509 14.0962C9.15254 13.1035 8.65338 12.6071 8.03422 12.6074C7.41506 12.6076 6.91626 13.1043 5.91867 14.0977L2 18"
                                    stroke="currentColor"
                                    strokeWidth="1.5"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                />
                                <path d="M22.0001 12.5458V7H16.418" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
                            </svg>
                        </div>
                        <div className="flex items-center justify-between">
                            <div className="w-full rounded-full h-5 p-1 bg-dark-light overflow-hidden shadow-3xl dark:shadow-none dark:bg-dark-light/10">
                                <div
                                    className="bg-gradient-to-r from-[#4361ee] to-[#805dca] w-full h-full rounded-full relative before:absolute before:inset-y-0 ltr:before:right-0.5 rtl:before:left-0.5 before:bg-white before:w-2 before:h-2 before:rounded-full before:m-auto"
                                    style={{ width: '65%' }}
                                ></div>
                            </div>
                            <span className="ltr:ml-5 rtl:mr-5 dark:text-white-light">57%</span>
                        </div>
                    </div>
                </div>
            </div>

            <div className="grid lg:grid-cols-3 gap-6 mb-6">
                <div className="panel h-full p-0 lg:col-span-2">
                    <div className="flex items-start justify-between dark:text-white-light mb-5 p-5 border-b  border-white-light dark:border-[#1b2e4b]">
                        <h5 className="font-semibold text-lg ">Call Logs</h5>
                        <div className="dropdown">
                            <Dropdown
                                offset={[0, 5]}
                                placement={`${isRtl ? 'bottom-start' : 'bottom-end'}`}
                                btnClassName="hover:text-primary"
                                button={
                                    <svg className="w-5 h-5 text-black/70 dark:text-white/70 hover:!text-primary" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <circle cx="5" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                        <circle opacity="0.5" cx="12" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                        <circle cx="19" cy="12" r="2" stroke="currentColor" strokeWidth="1.5" />
                                    </svg>
                                }
                            >
                                <ul>
                                    <li>
                                        <button type="button">View</button>
                                    </li>
                                    <li>
                                        <button type="button">Update</button>
                                    </li>
                                    <li>
                                        <button type="button">Delete</button>
                                    </li>
                                </ul>
                            </Dropdown>
                        </div>
                    </div>

                        <ReactApexChart options={uniqueVisitorSeries.options} series={uniqueVisitorSeries.series} type="bar" height={360} />
                </div>

                {/* <div
                        className="panel h-full overflow-hidden before:bg-[#1937cc] before:absolute before:-right-44 before:top-0 before:bottom-0 before:m-auto before:rounded-full before:w-96 before:h-96 grid grid-cols-1 content-between"
                        style={{ background: 'linear-gradient(0deg,#00c6fb -227%,#005bea)' }}
                    >
                        <div className="flex items-start justify-between text-white-light mb-16 z-[7]">

                            <div className="">
                                <svg className="svg-main-icon" id="Capa_1" enable-background="new 0 0 512 512" height="80" viewBox="0 0 512 512" width="80" xmlns="http://www.w3.org/2000/svg">
                                    <path d="m482 245.242v-60.363c0-29.656-23.597-53.891-53-54.949v-37.051c0-19.299-15.701-35-35-35h-96.358l-12.443-34.587c-3.173-8.82-9.595-15.868-18.083-19.845-8.488-3.978-18.014-4.402-26.821-1.196l-174.855 63.641c-8.798 3.202-15.817 9.641-19.765 18.131s-4.349 18.007-1.128 26.799l7.025 19.175c-28.735 1.777-51.572 25.707-51.572 54.882v272c0 30.327 24.673 55 55 55h372c30.327 0 55-24.673 55-55v-62.363c16.938-2.434 30-17.036 30-34.637v-80c0-17.601-13.062-32.203-30-34.637zm0 114.637c0 2.757-2.243 5-5 5h-80c-24.813 0-45-20.187-45-45s20.187-45 45-45h80c2.757 0 5 2.243 5 5zm-409.284-259.377c-.621-1.695-.166-3.126.161-3.829.327-.702 1.128-1.973 2.824-2.59l174.854-63.641c1.698-.617 3.129-.158 3.832.171s1.972 1.135 2.583 2.835l8.79 24.432h-6.76c-19.299 0-35 15.701-35 35v37h-140.521zm326.284-7.623v37h-145v-37c0-2.757 2.243-5 5-5h135c2.757 0 5 2.243 5 5zm28 389h-372c-13.785 0-25-11.215-25-25v-272c0-13.785 11.215-25 25-25h372c13.785 0 25 11.215 25 25v60h-55c-41.355 0-75 33.645-75 75s33.645 75 75 75h55v62c0 13.785-11.215 25-25 25z" fill="#fff"></path>
                                    <circle cx="397" cy="319.879" r="15" fill="#fff"></circle>
                                </svg>
                            </div>
                            <div className="font-semibold text-2xl">Total Balance</div>

                            <div className="relative text-xl whitespace-nowrap">
                                $ 41,741.42
                                <span className="table text-[#d3d3d3] bg-[#4361ee] rounded p-1 text-xs mt-1 ltr:ml-auto rtl:mr-auto">+ 2453</span>
                            </div>
                        </div>
                        <div className="flex items-center justify-between z-10">
                            <div className="flex items-center justify-between">
                                <button type="button" className="shadow-[0_0_2px_0_#bfc9d4] rounded p-1 text-white-light hover:bg-[#1937cc] place-content-center ltr:mr-2 rtl:ml-2">
                                    <svg className="w-5 h-5" viewBox="0 0 24 24" stroke="currentColor" strokeWidth="1.5" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                        <line x1="12" y1="5" x2="12" y2="19"></line>
                                        <line x1="5" y1="12" x2="19" y2="12"></line>
                                    </svg>
                                </button>
                                <button type="button" className="shadow-[0_0_2px_0_#bfc9d4] rounded p-1 text-white-light hover:bg-[#1937cc] grid place-content-center">
                                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path
                                            d="M2 12C2 8.22876 2 6.34315 3.17157 5.17157C4.34315 4 6.22876 4 10 4H14C17.7712 4 19.6569 4 20.8284 5.17157C22 6.34315 22 8.22876 22 12C22 15.7712 22 17.6569 20.8284 18.8284C19.6569 20 17.7712 20 14 20H10C6.22876 20 4.34315 20 3.17157 18.8284C2 17.6569 2 15.7712 2 12Z"
                                            stroke="currentColor"
                                            strokeWidth="1.5"
                                        />
                                        <path opacity="0.5" d="M10 16H6" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M14 16H12.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                        <path opacity="0.5" d="M2 10L22 10" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                                    </svg>
                                </button>
                            </div>
                            <button type="button" className="shadow-[0_0_2px_0_#bfc9d4] rounded p-1 text-white-light hover:bg-[#4361ee] z-10">
                                Upgrade
                            </button>
                        </div>
                </div> */}
            </div>
        </div>
    );
};

export default Home;