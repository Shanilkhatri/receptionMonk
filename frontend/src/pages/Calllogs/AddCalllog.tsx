import React from 'react';
import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';

const AddCalllog = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add Call Logs'));
    });

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Call Logs
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Add</span>
                </li>
            </ul>

            <div className='flex xl:flex-row flex-col gap-2.5 py-5'>
                <div className="panel px-0 flex-1 py-6 ltr:xl:mr-6 rtl:xl:ml-6">
                    <div className="mt-8 px-4">
                        <div className='text-xl font-bold text-dark dark:text-white text-center'>Add Call Logs</div>
                        <div className="flex justify-between lg:flex-row flex-col">
                            <div className="lg:w-1/2 w-full m-8">                           
                                <div className="my-8 flex items-center">
                                    <label htmlFor="call-from" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Call From
                                    </label>
                                    <input id="call-from" type="text" name="call-from" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Call From (e.g. Company, or Person's name)" />
                                </div>           
                                <div className="my-8 flex items-center">
                                    <label htmlFor="call-place" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Call Placed At
                                    </label>
                                    <input id="call-place" type="text" name="call-place" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Call Placed At (e.g., Office, Home)" />
                                </div>   
                                <div className="flex items-center my-8 ">
                                    <label htmlFor="call-duration" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Call Duration
                                    </label>
                                    <input id="call-duration" type="text" name="call-duration" className="form-input flex-1 border border-gray-400 focus:border-orange-400  text-gray-600" placeholder="Call Duration (e.g., Hours, Minutes and Seconds)" />
                                </div>                                  
                            </div>
                            <div className="lg:w-1/2 w-full m-8">               
                                <div className="my-8 flex items-center">
                                    <label htmlFor="call-to" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Call To
                                    </label>
                                    <input id="call-to" type="text" name="call-to" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Call To (e.g. Company, or Person's name)" />
                                </div>                                
                                <div className="my-8 flex items-center">
                                    <label htmlFor="call-ext" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Call Extension
                                    </label>
                                    <input id="call-ext" type="text" name="call-ext" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Call Extension (if applicable)" />
                                </div>     
                            </div>
                        </div>
                        <div className="flex w-full pb-12">
                            <div className="flex items-center lg:justify-end lg:mr-6 w-1/2">                        
                                <button type="button" className="btn btn-outline-dark rounded-xl px-6 font-bold">RESET</button>
                            </div>
                            <div className="flex items-center lg:ml-6 w-1/2">
                              
                                <button 
                                    type="submit" 
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:border-black font-bold" 
                                >
                                        ADD
                                </button>
                                
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AddCalllog;