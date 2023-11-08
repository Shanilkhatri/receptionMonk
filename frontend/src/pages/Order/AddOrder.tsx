import { Link , useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const AddOrder = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add Order'));
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
        title: 'Order Added Successfully',
        padding: '10px 20px',
    });
};

const SubmittedForm = Yup.object().shape({
    orderFrom: Yup.string().required('Please fill order purchase date'),
    orderTo: Yup.string().required('Please fill order expiry date'),
    orderAmount: Yup.string().required('Please fill order price'),
    orderBuyerName: Yup.string().required('Please fill buyer name'),
    userStatus: Yup.string().required('Please select order status'),   
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

            <Formik
                    initialValues={{
                        orderFrom: '',
                        orderTo: '',
                        orderAmount: '',
                        orderBuyerName: '',                        
                        userStatus: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            <div className='panel py-6 mt-6'>
                                <div className='p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8'>  
                                    <div className='text-xl font-bold text-dark dark:text-white text-center pb-12'>Add Order</div>      
                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="orderFrom" className='py-2'>
                                                    Order From <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.orderFrom ? 'has-error' : 'has-success') : ''}>
                                                <Field name="orderFrom" type="date" id="orderFrom" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter User Birth Date"  />

                                                {submitCount ? errors.orderFrom ? <div className="text-danger mt-1">{errors.orderFrom}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="orderTo" className='py-2'>
                                                    Order To <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.orderTo ? 'has-error' : 'has-success') : ''}>
                                                <Field name="orderTo" type="date" id="orderTo" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter User Birth Date"  />

                                                {submitCount ? errors.orderTo ? <div className="text-danger mt-1">{errors.orderTo}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="orderAmount" className='py-2'>
                                                    Price<span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.orderAmount ? 'has-error' : 'has-success') : ''}>
                                                <Field name="orderAmount" type="text" id="orderAmount" placeholder="Enter Order Amount" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.orderAmount ? <div className="text-danger mt-1">{errors.orderAmount}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="orderBuyerName" className="py-2">
                                                    Buyer <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.orderBuyerName ? 'has-error' : 'has-success') : ''}>
                                                <Field name="orderBuyerName" type="text" id="orderBuyerName" placeholder="Enter Buyer Name" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.orderBuyerName ? <div className="text-danger mt-1">{errors.orderBuyerName}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>                                            
                                        </div>                                     
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                       
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userStatus" className="py-2">
                                                    Status <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userStatus ? 'has-error' : 'has-success') : ''}>                                            

                                                <Field as="select" name="userStatus" className="form-select border border-gray-400 focus:border-orange-400 text-gray-600">
                                                    <option value="">Select</option>
                                                    <option value="name1">Paid</option>
                                                    <option value="name2">Unpaid</option>
                                                </Field>
                                                {submitCount ? (
                                                    errors.userStatus ? (
                                                        <div className=" text-danger mt-1">{errors.userStatus}</div>
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
                                                if (touched.orderFrom && !errors.orderFrom) {
                                                    submitForm();
                                                }
                                                else if (touched.orderTo && !errors.orderTo) {
                                                    submitForm();
                                                }
                                                else if (touched.orderAmount && !errors.orderAmount) {
                                                    submitForm();
                                                }
                                                else if (touched.orderBuyerName && !errors.orderBuyerName) {
                                                    submitForm();
                                                }
                                                else if (touched.userStatus && !errors.userStatus) {
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

export default AddOrder ;