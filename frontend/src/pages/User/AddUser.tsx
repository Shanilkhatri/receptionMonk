import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';

const AddUser = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add User'));
    });

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        User
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Add</span>
                </li>
            </ul>

            <div className='flex xl:flex-row flex-col gap-2.5 py-5'>
            <div className="panel px-0 flex-1 py-6 ltr:xl:mr-6 rtl:xl:ml-6">
                <div className="mt-8 px-4">
                    <div className='text-xl font-bold text-dark dark:text-white text-center'>Add User</div>
                    <div className="flex justify-between lg:flex-row flex-col">
                        <div className="lg:w-1/2 w-full m-8">                           
                            <div className="my-8 flex items-center">
                                <label htmlFor="reciever-name" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Name
                                </label>
                                <input id="reciever-name" type="text" name="reciever-name" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Name" />
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="reciever-email" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Email
                                </label>
                                <input id="reciever-email" type="email" name="reciever-email" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Email" />
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="reciever-address" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Password
                                </label>
                                <input id="reciever-address" type="text" name="reciever-address" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Password" />
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="reciever-number" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Phone Number
                                </label>
                                <input id="reciever-number" type="text" name="reciever-number" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Phone number" />
                            </div>
                        </div>
                        <div className="lg:w-1/2 w-full m-8">
                            <div className="flex items-center my-8">
                                <label htmlFor="acno" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Date of Birth
                                </label>
                                <input id="acno" type="date" name="acno" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Account Number" />
                            </div>
                            <div className="flex items-center my-8">
                                <label htmlFor="country" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Account Type
                                </label>
                                <select id="country" name="country" className="form-select flex-1 border border-gray-400 focus:border-orange-400">
                                    <option value="">Select</option>
                                    <option value="United States">User</option>
                                    <option value="United Kingdom">Owner</option>                   
                                </select>
                            </div>
                            <div className="flex items-center my-8">
                                <label htmlFor="swift-code" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Company Name
                                </label>
                                <input id="swift-code" type="text" name="swift-code" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Company Name" />
                            </div>                            
                            <div className="flex items-center my-8">
                                <label htmlFor="country" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Status
                                </label>
                                <select id="country" name="country" className="form-select flex-1 border border-gray-400 focus:border-orange-400">
                                    <option value="">Select</option>
                                    <option value="United States">Active</option>
                                    <option value="United Kingdom">Pending</option>
                                    <option value="Canada">Deactive</option>                                  
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
                    {/*  */}
                    {/* <div className="flex justify-center flex-col">
                        <div className="w-full mb-6">   
                            <div className="mt-4 flex items-center">
                                <label htmlFor="reciever-name" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Name
                                </label>
                                <input id="reciever-name" type="text" name="reciever-name" className="form-input flex-1" placeholder="Enter Name" />
                            </div>
                            <div className="mt-4 flex items-center">
                                <label htmlFor="reciever-email" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Email
                                </label>
                                <input id="reciever-email" type="email" name="reciever-email" className="form-input flex-1" placeholder="Enter Email" />
                            </div>
                            <div className="mt-4 flex items-center">
                                <label htmlFor="reciever-address" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Password
                                </label>
                                <input id="reciever-address" type="text" name="reciever-address" className="form-input flex-1" placeholder="Enter Password" />
                            </div>
                            <div className="mt-4 flex items-center">
                                <label htmlFor="reciever-number" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Phone Number
                                </label>
                                <input id="reciever-number" type="text" name="reciever-number" className="form-input flex-1" placeholder="Enter Phone number" />
                            </div>
                        </div>
                    </div> */}
                    {/*  */}
                </div>
                </div>
            </div>
        </div>
    );
};

export default AddUser;