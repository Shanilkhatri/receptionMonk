import { Link , useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const EditExt = () => {

const dispatch = useDispatch();
useEffect(() => {
    dispatch(setPageTitle('Add Extension'));
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
        title: 'Extenstion Updated Successfully',
        padding: '10px 20px',
    });
};

const SubmittedForm = Yup.object().shape({
    extName: Yup.string().required('Please fill Extension Name'),
    extUserName: Yup.string().required('Please select Extention User'),
    extDept: Yup.string().required('Please fill Department'),
    extSipServer: Yup.string().required('Please fill Extension Sip Server'),
    extSipUserName: Yup.string().required('Please fill Extension Sip User Name'),
    extSipPwd: Yup.string().required('Please fill Extension Sip Password'),
    extSipPort: Yup.string().required('Please fill Extension Sip Port'),
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

            <Formik
                    initialValues={{
                        extName: '',
                        extUserName: '',
                        extDept: '',
                        extSipServer: '',
                        extSipUserName: '', 
                        extSipPwd: '',
                        extSipPort: '',                       
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            <div className='panel py-6 mt-6'>
                                <div className='p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8'>  
                                    <div className='text-xl font-bold text-dark dark:text-white text-center pb-12'>Update Extension</div>      
                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="extName" className='py-2'>
                                                    Extension Name <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extName ? 'has-error' : 'has-success') : ''}>
                                                <Field name="extName" type="text" id="extName" placeholder="Enter Extension Name" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.extName ? <div className="text-danger mt-1">{errors.extName}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="extUserName" className="py-2">
                                                    User Name <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extUserName ? 'has-error' : 'has-success') : ''}>
                                                <Field as="select" name="extUserName" className="form-select border border-gray-400 focus:border-orange-400 text-gray-600">
                                                    <option value="">Select</option>
                                                    <option value="name1">Name 1</option>
                                                    <option value="name2">Name 2</option>            
                                                </Field>
                                                {submitCount ? (
                                                    errors.extUserName ? (
                                                        <div className=" text-danger mt-1">{errors.extUserName}</div>
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
                                                <label htmlFor="extDept" className='py-2'>
                                                    Department <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extDept ? 'has-error' : 'has-success') : ''}>
                                                <Field name="extDept" type="text" id="extDept" placeholder="Enter Department Name" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.extDept ? <div className="text-danger mt-1">{errors.extDept}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="extSipServer" className="py-2">
                                                    SIP Server <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extSipServer ? 'has-error' : 'has-success') : ''}>
                                                <Field name="extSipServer" type="text" id="extSipServer" placeholder="Enter Sip Server" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.extSipServer ? <div className="text-danger mt-1">{errors.extSipServer}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>                                            
                                        </div>                                     
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="extSipUserName" className='py-2'>
                                                    SIP Username <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extSipUserName ? 'has-error' : 'has-success') : ''}>
                                                <Field name="extSipUserName" type="text" id="extSipUserName" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter Sip Username"  />

                                                {submitCount ? errors.extSipUserName ? <div className="text-danger mt-1">{errors.extSipUserName}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="extSipPwd" className="py-2">
                                                    SIP Password <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extSipPwd ? 'has-error' : 'has-success') : ''}>                                            

                                                <Field name="extSipPwd" type="text" id="extSipPwd" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter Sip Password"  />

                                                {submitCount ? errors.extSipPwd ? <div className="text-danger mt-1">{errors.extSipPwd}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                               
                                            </div>                                            
                                        </div>
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="extSipPort" className='py-2'>
                                                    SIP Port <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.extSipPort ? 'has-error' : 'has-success') : ''}>
                                                <Field name="extSipPort" type="text" id="extSipPort" placeholder="Enter Sip Port" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.extSipPort ? <div className="text-danger mt-1">{errors.extSipPort}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
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
                                                if (touched.extName && !errors.extName) {
                                                    submitForm();
                                                }                                                
                                                else if (touched.extUserName && !errors.extUserName) {
                                                    submitForm();
                                                }
                                                else if (touched.extDept && !errors.extDept) {
                                                    submitForm();
                                                }
                                                else if (touched.extSipServer && !errors.extSipServer) {
                                                    submitForm();
                                                }
                                                else if (touched.extSipUserName && !errors.extSipUserName) {
                                                    submitForm();
                                                }
                                                else if (touched.extSipPwd && !errors.extSipPwd) {
                                                    submitForm();
                                                } 
                                                else if (touched.extSipPort && !errors.extSipPort) {
                                                    submitForm();
                                                }
                                            }}
                                        >
                                                UPDATE
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

export default EditExt;