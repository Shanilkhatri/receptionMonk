import React from "react";

const Submitform = () => {
  return (
    <div>
      <Form className="space-y-5">
        <div className="panel py-6 mt-6">
          <div className="p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8">
            <div className="text-xl font-bold text-dark dark:text-white text-center pb-12">
              Add User
            </div>
            <div className="grid lg:grid-cols-2 lg:space-x-12 lg:my-5">
              <div className="grid md:grid-cols-2 my-3 lg:my-0">
                <div className="">
                  <label htmlFor="userName" className="py-2">
                    Name <span className="text-red-600">*</span>
                  </label>
                </div>
                <div
                  className={
                    submitCount
                      ? errors.userName
                        ? "has-error"
                        : "has-success"
                      : ""
                  }
                >
                  <Field
                    name="userName"
                    type="text"
                    id="userName"
                    placeholder="Enter User Name"
                    className="form-input border border-gray-400 focus:border-orange-400"
                  />

                  {submitCount ? (
                    errors.userName ? (
                      <div className="text-danger mt-1">{errors.userName}</div>
                    ) : (
                      <div className="text-success mt-1">It is fine!</div>
                    )
                  ) : (
                    ""
                  )}
                </div>
              </div>

              <div className="grid md:grid-cols-2 my-3 lg:my-0">
                <div className="">
                  <label htmlFor="user-email" className="py-2">
                    Email <span className="text-red-600">*</span>
                  </label>
                </div>
                <div
                  className={
                    submitCount
                      ? errors.userEmail
                        ? "has-error"
                        : "has-success"
                      : ""
                  }
                >
                  <Field
                    name="userEmail"
                    type="text"
                    id="userEmail"
                    placeholder="Enter User Email"
                    className="form-input border border-gray-400 focus:border-orange-400"
                  />

                  {submitCount ? (
                    errors.userEmail ? (
                      <div className="text-danger mt-1">{errors.userEmail}</div>
                    ) : (
                      <div className="text-success mt-1">It is fine!</div>
                    )
                  ) : (
                    ""
                  )}
                </div>
              </div>
            </div>

            <div className="grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5">
              <div className="grid md:grid-cols-2 my-3 lg:my-0">
                <div className="">
                  <label htmlFor="userDob" className="py-2">
                    Date of Birth <span className="text-red-600">*</span>
                  </label>
                </div>
                <div
                  className={
                    submitCount
                      ? errors.userDob
                        ? "has-error"
                        : "has-success"
                      : ""
                  }
                >
                  <Field
                    name="userDob"
                    type="date"
                    id="userDob"
                    className="form-input border border-gray-400 focus:border-orange-400 text-gray-600"
                    placeholder="Enter User Birth Date"
                  />

                  {submitCount ? (
                    errors.userDob ? (
                      <div className="text-danger mt-1">{errors.userDob}</div>
                    ) : (
                      <div className="text-success mt-1">It is fine!</div>
                    )
                  ) : (
                    ""
                  )}
                </div>
              </div>

              <div className="grid  md:grid-cols-2 my-3 lg:my-0">
                <div className="">
                  <label htmlFor="userAccType" className="py-2">
                    Account Type <span className="text-red-600">*</span>
                  </label>
                </div>
                <div
                  className={
                    submitCount
                      ? errors.userAccType
                        ? "has-error"
                        : "has-success"
                      : ""
                  }
                >
                  <Field
                    as="select"
                    name="userAccType"
                    id="userAccType"
                    className="form-select border border-gray-400 focus:border-orange-400 text-gray-600"
                  >
                    <option value="">Select</option>
                    <option value="User">User</option>
                    <option value="Owner">Owner</option>
                  </Field>
                  {submitCount ? (
                    errors.userAccType ? (
                      <div className=" text-danger mt-1">
                        {errors.userAccType}
                      </div>
                    ) : (
                      <div className="text-success mt-1">It is fine!</div>
                    )
                  ) : (
                    ""
                  )}
                </div>
              </div>
            </div>

            <div className="grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5">
              {/*      <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userCompName" className='py-2'>
                                                    Company Name <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userCompName ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userCompName" type="text" id="userCompName" placeholder="Enter Company Name" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.userCompName ? <div className="text-danger mt-1">{errors.userCompName}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div> */}

              <div className="grid  md:grid-cols-2 my-3 lg:my-0">
                <div className="">
                  <label htmlFor="userStatus" className="py-2">
                    Status <span className="text-red-600">*</span>
                  </label>
                </div>
                <div
                  className={
                    submitCount
                      ? errors.userStatus
                        ? "has-error"
                        : "has-success"
                      : ""
                  }
                >
                  <Field
                    as="select"
                    name="userStatus"
                    id="userStatus"
                    className="form-select border border-gray-400 focus:border-orange-400 text-gray-600"
                  >
                    <option value="">Select</option>
                    <option value="United States">Active</option>
                    <option value="United Kingdom">Pending</option>
                    <option value="Canada">Deactive</option>
                  </Field>
                  {submitCount ? (
                    errors.userStatus ? (
                      <div className=" text-danger mt-1">
                        {errors.userStatus}
                      </div>
                    ) : (
                      <div className="text-success mt-1">It is fine!</div>
                    )
                  ) : (
                    ""
                  )}
                </div>
              </div>

              {/* avatar added */}
              <div className="grid md:grid-cols-2 my-3 lg:my-0">
                <div className="">
                  <label htmlFor="user-avatar" className="py-2">
                    Avatar <span className="text-red-600">*</span>
                  </label>
                </div>
                <div
                  className={
                    submitCount
                      ? errors.avatar
                        ? "has-error"
                        : "has-success"
                      : ""
                  }
                >
                  <Foravatar />
                  {/* <Field
                        name="userAvatar"
                        type="image"
                        id="userAvatar"
                        placeholder="Browse"
                        className="form-input border border-gray-400 focus:border-orange-400"
                      /> */}

                  {submitCount ? (
                    errors.avatar ? (
                      <div className="text-danger mt-1">{errors.avatar}</div>
                    ) : (
                      <div className="text-success mt-1">It is fine!</div>
                    )
                  ) : (
                    ""
                  )}
                </div>
                {/* <div>
                      {" "}
                      <input type="checkbox" /> Default value
                    </div> */}
              </div>
            </div>
            <div className="flex justify-center py-6 mt-12">
              <button
                type="reset"
                className="btn btn-outline-dark rounded-xl px-6 mx-3 font-bold"
              >
                RESET
              </button>

              <button
                type="submit"
                className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828] mx-3"
                onClick={() => {
                  if (touched.userName && !errors.userName) {
                    submitForm();
                  } else if (touched.userEmail && !errors.userEmail) {
                    submitForm();
                  }
                  // else if (touched.userPassword && !errors.userPassword) {
                  //     submitForm();
                  // }
                  // else if (touched.userPhone && !errors.userPhone) {
                  //     submitForm();
                  // }
                  else if (touched.userDob && !errors.userDob) {
                    submitForm();
                  } else if (touched.userAccType && !errors.userAccType) {
                    submitForm();
                  }
                  // else if (touched.userCompName && !errors.userCompName) {
                  //     submitForm();
                  // }
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
    </div>
  );
};

export default Submitform;
