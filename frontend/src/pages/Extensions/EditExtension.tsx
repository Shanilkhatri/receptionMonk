import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';

const EditExt = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Update Extension'));
    });

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Extension
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Update</span>
                </li>
            </ul>

            <div className='flex xl:flex-row flex-col gap-2.5 py-5'>
                <div className="panel px-0 flex-1 py-6 ltr:xl:mr-6 rtl:xl:ml-6">
                <div className="mt-8 px-4">
                    <div className='text-xl font-bold text-dark dark:text-white text-center'>Update Extension</div>
                    <div className="flex justify-between lg:flex-row flex-col">
                        <div className="lg:w-1/2 w-full m-8">                           
                            <div className="my-8 flex items-center">
                                <label htmlFor="ext-name" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                   Extension Name
                                </label>
                                <input id="ext-name" type="text" name="ext-name" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Name" />
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="username" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                   User Name
                                </label>
                                <select id="username" name="username" className="form-select flex-1 border border-gray-400 focus:border-orange-400">    
                                    <option value="">Select</option>
                                    <option value="name1">Name 1</option>
                                    <option value="name2">Name 2</option>                   
                                </select>
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="dept-name" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    Department
                                </label>
                                <input id="dept-name" type="text" name="dept-name" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Department Name" />
                            </div>
                            <div className="my-8 flex items-center">
                                <label htmlFor="sip-server" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    SIP Server
                                </label>
                                <input id="sip-server" type="text" name="sip-server" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Sip Server" />
                            </div>
                        </div>
                        <div className="lg:w-1/2 w-full m-8">
                            <div className="flex items-center my-8">
                                <label htmlFor="sip-user" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    SIP Username
                                </label>
                                <input id="sip-user" type="text" name="sip-user" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Sip Username" />
                            </div>
                            <div className="flex items-center my-8">
                                <label htmlFor="sip-pwd" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    SIP Password
                                </label>
                                <input id="sip-pwd" type="sip-pwd" name="acno" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Sip Password" />
                            </div>
                            <div className="flex items-center my-8">
                                <label htmlFor="sip-port" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                    SIP Port
                                </label>
                                <input id="sip-port" type="text" name="sip-port" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Sip Port" />
                            </div>       
                        </div>
                    </div>
                    <div className="flex w-full pb-12">
                        <div className="flex items-center lg:justify-end lg:mr-6 w-1/2">                        
                            <button type="reset" className="btn btn-outline-dark rounded-xl px-6 font-bold">RESET</button>
                        </div>
                        <div className="flex items-center lg:ml-6 w-1/2">
                            <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-6 hover:border-black font-bold">UPDATE</button>
                        </div>
                    </div>                    
                </div>
                </div>
            </div>
        </div>
    );
};

export default EditExt;