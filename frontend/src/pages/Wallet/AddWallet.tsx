import { Link , useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const AddWallet = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add Wallet'));
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
        title: 'Wallet Added Successfully',
        padding: '10px 20px',
    });
};

const SubmittedForm = Yup.object().shape({
    walletPlan: Yup.string().required('Please select credit plan'),
    compName: Yup.string().required('Please fill company name'),
    walletAmount: Yup.string().required('Please fill wallet amount'),
    walletDate: Yup.string().required('Please pick a date'),
    reason: Yup.string().required('Please fill reason'),   
});

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Wallet
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Add</span>
                </li>
            </ul>

            <Formik
                initialValues={{
                    walletPlan: '',
                    walletAmount: '',
                    walletDate: '',
                    compName: '',                        
                    reason: '',
                }}
                validationSchema={SubmittedForm}
                onSubmit={() => {}}
            >
                {({ errors, submitCount, touched })  => (   
                    <Form className="space-y-5">
                        <div className='panel py-6 mt-6'>
                            <div className='p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8'>  
                                <div className='text-xl font-bold text-dark dark:text-white text-center pb-12'>Add Wallet</div>      
                                <div className='grid lg:grid-cols-2 lg:space-x-12 lg:my-5'>
                                    <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                        <div className=''>
                                            <label htmlFor="walletPlan" className='py-2'>
                                                Charge Plan <span className='text-red-600'>*</span>
                                            </label>
                                        </div>
                                        <div className={submitCount ? (errors.walletPlan ? 'has-error' : 'has-success') : ''}>
                                            <Field as="select" name="walletPlan" className="form-select border border-gray-400 focus:border-orange-400 text-gray-600">
                                            <option value="">Select Credit</option>
                                            <option value="United States">Plan 1</option>
                                            <option value="United Kingdom">Plan 2</option>                   
                                            </Field>
                                            {submitCount ? (
                                                errors.walletPlan ? (
                                                    <div className=" text-danger mt-1">{errors.walletPlan}</div>
                                                ) : (
                                                    <div className="text-success mt-1">It is fine!</div>
                                                )
                                            ) : (
                                                ''
                                            )}                                            
                                        </div>
                                    </div>
                                    
                                    <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                        <div className=''>
                                            <label htmlFor="compName" className='py-2'>
                                                Company Name <span className='text-red-600'>*</span>
                                            </label>
                                        </div>
                                        <div className={submitCount ? (errors.compName ? 'has-error' : 'has-success') : ''}>
                                            <Field as="select" name="compName" className="form-select border border-gray-400 focus:border-orange-400 text-gray-600">
                                                <option value="">Select</option>
                                                <option value="United States">HP</option>
                                                <option value="United Kingdom">Dell</option>                                   
                                            </Field>
                                            {submitCount ? (
                                                errors.compName ? (
                                                    <div className=" text-danger mt-1">{errors.compName}</div>
                                                ) : (
                                                    <div className="text-success mt-1">It is fine!</div>
                                                )
                                            ) : (
                                                ''
                                            )}
                                        </div>
                                    </div>
                                </div>

                                <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                    <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                        <div className=''>
                                            <label htmlFor="walletAmount" className='py-2'>
                                                Amount<span className='text-red-600'>*</span>
                                            </label>
                                        </div>
                                        <div className={submitCount ? (errors.walletAmount ? 'has-error' : 'has-success') : ''}>
                                            <Field name="walletAmount" type="text" id="walletAmount" placeholder="Enter Order Amount" className="form-input border border-gray-400 focus:border-orange-400" />

                                            {submitCount ? errors.walletAmount ? <div className="text-danger mt-1">{errors.walletAmount}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                        </div>
                                    </div>
                                    
                                    <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                        <div className=''>
                                            <label htmlFor="walletDate" className="py-2">
                                                Date <span className='text-red-600'>*</span>
                                            </label>
                                        </div>
                                        <div className={submitCount ? (errors.walletDate ? 'has-error' : 'has-success') : ''}>
                                            <Field name="walletDate" type="date" id="walletDate" placeholder="Enter Date" className="form-input border border-gray-400 focus:border-orange-400" />

                                            {submitCount ? errors.walletDate ? <div className="text-danger mt-1">{errors.walletDate}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                        </div>                                            
                                    </div>                                     
                                </div>

                                <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                    
                                    <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                        <div className=''>
                                            <label htmlFor="reason" className="py-2">
                                                Reason <span className='text-red-600'>*</span>
                                            </label>
                                        </div>
                                        <div className={submitCount ? (errors.reason ? 'has-error' : 'has-success') : ''}>                                            

                                            <Field as="textarea" name="reason" className="form-textarea border border-gray-400 focus:border-orange-400 flex-1" placeholder="Enter Description"></Field>
                                            {submitCount ? (
                                                errors.reason ? (
                                                    <div className=" text-danger mt-1">{errors.reason}</div>
                                                ) : (
                                                    <div className="text-success mt-1">It is fine!</div>
                                                )
                                            ) : (
                                                ''
                                            )}
                                            
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
                                            if (touched.walletPlan && !errors.walletPlan) {
                                                submitForm();
                                            }
                                            else if (touched.compName && !errors.compName) {
                                                submitForm();
                                            }
                                            else if (touched.walletAmount && !errors.walletAmount) {
                                                submitForm();
                                            }
                                            else if (touched.walletDate && !errors.walletDate) {
                                                submitForm();
                                            }
                                            else if (touched.reason && !errors.reason) {
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

export default AddWallet;