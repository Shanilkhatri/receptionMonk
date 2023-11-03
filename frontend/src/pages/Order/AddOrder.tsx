import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';

const AddOrder = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add Order'));
    });

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Order
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Add</span>
                </li>
            </ul>

            <div className='flex xl:flex-row flex-col gap-2.5 py-5'>
                <div className="panel px-0 flex-1 py-6 ltr:xl:mr-6 rtl:xl:ml-6">
                <div className="mt-8 px-4">
                    <div className='text-xl font-bold text-dark dark:text-white text-center'>Add Order</div>
                    <div className="flex justify-between lg:flex-row flex-col">
                        <div className="lg:w-1/2 w-full m-8">                           
                            <div className="my-8 flex items-center">
                                <label htmlFor="order-from-date" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                   Order From
                                </label>
                                <input id="order-from-date" type="date" name="order-from-date" className="form-input flex-1 border border-gray-400 focus:border-orange-400 text-gray-500" />
                            </div>                                                     
                            <div className="my-8 flex items-center">
                                <label htmlFor="order-price" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                   Price
                                </label>
                                <input id="order-price" type="text" name="order-price" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Order Amount" />
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="order-buyer" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Buyer
                                </label>
                                <input id="order-buyer" type="text" name="order-buyer" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Buyer Name" />
                            </div>
                        </div>
                        <div className="lg:w-1/2 w-full m-8">    
                            <div className="my-8 flex items-center">
                                <label htmlFor="order-to-date" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                   Order Upto
                                </label>
                                <input id="order-to-date" type="date" name="order-to-date" className="form-input flex-1 border border-gray-400 focus:border-orange-400 text-gray-500"/>
                            </div>                          
                            <div className="my-8 flex items-center">
                                <label htmlFor="username" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                   Status
                                </label>
                                <select id="username" name="username" className="form-select flex-1 border border-gray-400 focus:border-orange-400  text-gray-500">    
                                    <option value="">Select</option>
                                    <option value="name1">Paid</option>
                                    <option value="name2">Unpaid</option>                   
                                </select>
                            </div>
                        </div>
                    </div>
                    <div className="flex w-full pb-12">
                        <div className="flex items-center lg:justify-end lg:mr-6 w-1/2">                        
                            <button type="button" className="btn btn-outline-dark rounded-xl px-6 font-bold">RESET</button>
                        </div>
                        <div className="flex items-center lg:ml-6 w-1/2">
                            <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:border-black font-bold">ADD</button>
                        </div>
                    </div>                    
                </div>
                </div>
            </div>
        </div>
    );
};

export default AddOrder ;