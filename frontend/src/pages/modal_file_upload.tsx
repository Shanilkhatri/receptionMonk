// import React, { Component, ChangeEvent, FormEvent } from 'react';
// import { connect } from 'react-redux';
// import axios, { AxiosResponse } from 'axios';
// import { Dispatch } from 'redux';
// import { RootState } from './path-to-your-root-reducer'; // Update the path accordingly

// interface FileUploadProps {
//     file: File;
//     error: boolean;
//     showElement: boolean;
//     uploaded: number | null;
//     isLoading: (loading: boolean) => void;
//     dispatch: Dispatch;
//     errorOrders: (payload: any) => void; // Adjust the type of payload as per your implementation
// }

// class FileUpload extends Component<FileUploadProps> {
//     fileuploaded(event: FormEvent) {
//         event.preventDefault();
//         const form = new FormData();
//         form.append('file', this.props.file);

//         if (this.props.file && this.props.file.name !== '') {
//             this.props.isLoading(true);

//             axios.post('http://localhost:4000/kycfileupload', form, {
//                 headers: { 'Content-Type': 'multipart/form-data' },
//                 onUploadProgress: (data:any) => {
//                     this.props.dispatch({ type: 'SET_UPLOADED', payload: Math.round((data.loaded / data.total) * 100) });
//                 },
//             })
//                 .then((response: AxiosResponse) => response.data)
//                 .then((response:any) => {
//                     this.props.isLoading(false);

//                     if (response.Payload.length > 0) {
//                         this.props.dispatch({ type: 'SET_SHOW_ELEMENT', payload: false });
//                         alert(response.Message);
//                         this.props.errorOrders(response.Payload);
//                     } else if (response.Payload.length === 0) {
//                         (event.target as HTMLFormElement).reset();
//                         alert(response.Message);
//                         this.props.dispatch({ type: 'SET_UPLOADED', payload: null });
//                     }
//                 })
//                 .catch((error:Error) => {
//                     console.error('Error:', error);
//                     alert(error);
//                 });
//         } else {
//             this.props.dispatch({ type: 'SET_ERROR', payload: true });
//         }
//     }

//     handleFileChange(event: ChangeEvent<HTMLInputElement>) {
//         if (event.target.files) {
//             this.props.dispatch({ type: 'SET_FILE', payload: event.target.files[0] });
//         }
//     }

//     render() {
//         return (
//             <>
//                 {this.props.showElement && (
//                     <div>
//                         <div className='row'>
//                             <form encType='multipart/form-data' onSubmit={(event) => this.fileuploaded(event)}>
//                                 <input type='file' placeholder='select file' onChange={(event) => this.handleFileChange(event)} />
//                                 {this.props.error ? <label>select excel file only</label> : ''}
//                                 <button type='submit'>upload</button>
//                             </form>
//                         </div>
//                         {this.props.uploaded && (
//                             <div className='progress mt-2 row' style={{ width: 350 }}>
//                                 <div
//                                     className='progress-bar'
//                                     role='progressbar'
//                                     aria-valuenow={this.props.uploaded}
//                                     aria-valuemin={0}
//                                     aria-valuemax={100}
//                                     style={{ width: `${this.props.uploaded}%` }}
//                                 >
//                                     {`${this.props.uploaded}%`}
//                                 </div>
//                             </div>
//                         )}
//                     </div>
//                 )}
//             </>
//         );
//     }
// }

// const mapStateToProps = (state: RootState) => ({
//     file: state.file,
//     error: state.error,
//     showElement: state.showElement,
//     uploaded: state.uploaded,
// });

// export default connect(mapStateToProps)(FileUpload);
