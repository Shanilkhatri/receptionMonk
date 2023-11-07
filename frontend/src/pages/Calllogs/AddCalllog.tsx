import { Link , useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const AddCalllog = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add Call Logs'));
    });

const navigate = useNavigate();

const submitForm = () => {
    // navigate('/');
    const toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
    });
    toast.fire({
        icon: 'success',
        title: 'Call Logs Added Successfully',
        padding: '10px 20px',
    });
};

const SubmittedForm = Yup.object().shape({
    callFrom: Yup.string().required('Please fill call from detail'),
    callTo: Yup.string().required('Please fill call to detail'),
    callPlacedAt: Yup.string().required('Please fill call place detail'),
    callDuration: Yup.string().required('Please fill call duration'),
    callExtension: Yup.string().required('Please fill call extension (if applicable)'),   
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

            <Formik
                    initialValues={{
                        callFrom: '',
                        callTo: '',
                        callPlacedAt: '',
                        callDuration: '',                        
                        callExtension: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            <div className='panel py-6 mt-6'>
                                <div className='p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8'>  
                                    <div className='text-xl font-bold text-dark dark:text-white text-center pb-12'>Add Call Logs</div>      
                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="callFrom" className='py-2'>
                                                    Call From <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.callFrom ? 'has-error' : 'has-success') : ''}>
                                                <Field name="callFrom" type="date" id="callFrom" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter Call From (e.g. Company, or Person's name)"  />

                                                {submitCount ? errors.callFrom ? <div className="text-danger mt-1">{errors.callFrom}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="callTo" className='py-2'>
                                                    Call To <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.callTo ? 'has-error' : 'has-success') : ''}>
                                                <Field name="callTo" type="date" id="callTo" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter Call To (e.g. Company, or Person's name)"  />

                                                {submitCount ? errors.callTo ? <div className="text-danger mt-1">{errors.callTo}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="callPlacedAt" className='py-2'>
                                                    Call Placed At<span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.callPlacedAt ? 'has-error' : 'has-success') : ''}>
                                                <Field name="callPlacedAt" type="text" id="callPlacedAt" placeholder="Enter Call Placed At (e.g., Office, Home)" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.callPlacedAt ? <div className="text-danger mt-1">{errors.callPlacedAt}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="callDuration" className="py-2">
                                                    Call Duration <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.callDuration ? 'has-error' : 'has-success') : ''}>
                                                <Field name="callDuration" type="text" id="callDuration" placeholder="Call Duration (e.g., Hours, Minutes and Seconds)" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.callDuration ? <div className="text-danger mt-1">{errors.callDuration}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>                                            
                                        </div>                                     
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                       
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="callExtension" className="py-2">
                                                    Call Extension <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.callExtension ? 'has-error' : 'has-success') : ''}>                                            

                                            <Field name="callExtension" type="text" id="callExtension" placeholder="Enter Call Extension (if applicable)" className="form-input border border-gray-400 focus:border-orange-400" />

                                            {submitCount ? errors.callExtension ? <div className="text-danger mt-1">{errors.callExtension}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                               
                                            </div>                                            
                                        </div>
                                    </div>

                                    <div className="flex justify-center py-6 mt-12">
                                        <button 
                                            type="reset" 
                                            className="btn btn-outline-dark rounded-xl px-6 mx-3 font-bold">
                                            RESET
                                        </button>

                                        <button
                                            type="submit"
                                            className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828] mx-3"
                                            onClick={() => {
                                                if (touched.callFrom && !errors.callFrom) {
                                                    submitForm();
                                                }
                                                else if (touched.callTo && !errors.callTo) {
                                                    submitForm();
                                                }
                                                else if (touched.callPlacedAt && !errors.callPlacedAt) {
                                                    submitForm();
                                                }
                                                else if (touched.callDuration && !errors.callDuration) {
                                                    submitForm();
                                                }
                                                else if (touched.callExtension && !errors.callExtension) {
                                                    submitForm();
                                                }
                                            }}
                                        >
                                                ADD
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </Form>
                     )}
            </Formik> 
        </div>
    );
};

export default AddCalllog;