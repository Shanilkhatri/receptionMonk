import { Link, useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import store, { IRootState } from "../../store";
import { setPageTitle } from "../../store/themeConfigSlice";
import * as Yup from "yup";
import { useEffect, useState } from "react";
import { Field, Form, Formik, useFormikContext } from "formik";
import Swal from "sweetalert2";
import Utility from "../../utility/utility";
const utility = new Utility();

const EditUser = () => {
  const [image, setImage] = useState<string | undefined>();
  // const formik = useFormikContext();

  // const isImageSet = useSelector((state) => state.image.isImageSet); //checking for image set or not

  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(setPageTitle("Update User"));
    if (store.getState().themeConfig.currentUserDataForUpdate.email == "") {
      navigate("/ViewUser");
    }
    setImage(
      import.meta.env.VITE_APPURL +
        store.getState().themeConfig.currentUserDataForUpdate.avatar
    );
  }, []);

  const navigate = useNavigate();

  function handleChange(e: any) {
    e.preventDefault();
    //const file = e.target.files[0];

    // console.log("file: ", URL.createObjectURL(e.target.files[0]));
    setImage(URL.createObjectURL(e.target.files[0]));
  }

  async function editUser(data: any) {
    let img = await utility.uploadImageAndReturnUrl(
      "fileInput",
      "avatar",
      "kycfileupload"
    );
    let userId = store.getState().themeConfig.currentUserDataForUpdate.id;
    if (typeof userId == "string") {
      userId = parseInt(userId);
    }
    let companyId =
      store.getState().themeConfig.currentUserDataForUpdate.companyId;
    if (typeof companyId == "string") {
      companyId = parseInt(companyId);
    }

    const userData = {
      id: userId,
      name: data.userName,
      email: data.userEmail,
      dob: data.userDob,
      status: data.userStatus,
      accountType: data.userAccType,
      avatar: img,
      companyId: companyId,
      iswizardcomplete: "kyc",
    };
    const ok = await utility.sendRequest_Put_Post_Get(userData, "users", "POST");
    if (ok.Status == "200") {
      navigate("/viewuser");
      return;
    } 
      return;
  }
  const schema = Yup.object().shape(
    {
      userName: Yup.string().required("Please fill User Name"),
      userEmail: Yup.string()
        .email("Invalid Email Address")
        .required("Please fill Email"),
      // userDob: Yup.date().required("Please enter a valid date."),
      userDob: Yup.date()
        .required("Please enter a valid date.")
        .min(
          new Date(new Date().setFullYear(new Date().getFullYear() - 70)),
          "Must be at most 70 years old"
        )
        .max(
          new Date(new Date().setFullYear(new Date().getFullYear() - 18)),
          "Must be at least 18 years old"
        ),
      userAccType: Yup.string().required("Please select User Type"),
      userStatus: Yup.string().required("Please select User Status"),
      avatar: Yup.mixed().when("avatar", {
        is: (value: any) => value?.length,
        then: (schema) =>
          schema.test("avatar", "image is required", (value) => {
            return value != undefined && value[0] && value[0].avatar !== "";
          }),
        otherwise: (schema) => schema.nullable(),
      }),
    },
    //cyclic dependencies

    [["avatar", "avatar"]]
  );

  const isRtl =
    useSelector((state: IRootState) => state.themeConfig.rtlClass) === "rtl"
      ? true
      : false;

  return (
    <div>
      <ul className="flex space-x-2 rtl:space-x-reverse">
        <li>
          <Link to="#" className="text-primary hover:underline">
            User
          </Link>
        </li>
        <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
          <span>Update</span>
        </li>
      </ul>

      <Formik
        initialValues={{
          id: store.getState().themeConfig.currentUserDataForUpdate.id,
          userName: store.getState().themeConfig.currentUserDataForUpdate.name,
          userEmail:
            store.getState().themeConfig.currentUserDataForUpdate.email,
          userDob: store.getState().themeConfig.currentUserDataForUpdate.dob,
          userAccType:
            store.getState().themeConfig.currentUserDataForUpdate.accountType,
          userStatus:
            store.getState().themeConfig.currentUserDataForUpdate.status,
          avatar:
            import.meta.env.VITE_APPURL +
            store.getState().themeConfig.currentUserDataForUpdate.avatar,
        }}
        validationSchema={schema}
        //    Yup.object().shape({
        //   userName: Yup.string().required("Please fill User Name"),
        //   userEmail: Yup.string()
        //     .email("Invalid Email Address")
        //     .required("Please fill Email"),
        //   userDob: Yup.date().required("Please enter a valid date."),
        //   userAccType: Yup.string().required("Please select User Type"),
        //   userStatus: Yup.string().required("Please select User Status"),
        //   avatar:Yup.mixed()
        // })

        onSubmit={(values, { setSubmitting }) => {
          setTimeout(() => {
            console.log(values);
            editUser(values);
            setSubmitting(false);
          }, 400);
        }}
      >
        {({ errors, submitCount, touched }) => (
          <Form className="space-y-5">
            <div className="panel py-6 mt-6">
              <div className="p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8">
                <div className="text-xl font-bold text-dark dark:text-white text-center pb-12">
                  Update User
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
                          <div className="text-danger mt-1">
                            {errors.userName}
                          </div>
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
                          <div className="text-danger mt-1">
                            {errors.userEmail}
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

                <div className="grid lg:grid-cols-2 lg:space-x-12 lg:my-5">
                  <div className="grid md:grid-cols-2 my-3 lg:my-0">
                    <div className="">
                      <label htmlFor="id" className="py-2">
                        id <span className="text-red-600">*</span>
                      </label>
                    </div>
                    <div
                      className={
                        submitCount
                          ? errors.id
                            ? "has-error"
                            : "has-success"
                          : ""
                      }
                    >
                      <Field
                        name="id"
                        type="text"
                        id="id"
                        placeholder="Enter User Name"
                        className="form-input border border-gray-400 focus:border-orange-400"
                      />

                      {submitCount ? (
                        errors.id ? (
                          <div className="text-danger mt-1">{errors.id}</div>
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
                          <div className="text-danger mt-1">
                            {errors.userDob}
                          </div>
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
                        className="form-select border border-gray-400 focus:border-orange-400 text-gray-600"
                      >
                        <option value="">Select</option>
                        <option value="user">User</option>
                        <option value="owner">Owner</option>
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
                        className="form-select border border-gray-400 focus:border-orange-400 text-gray-600"
                      >
                        <option value="">Select</option>
                        <option value="active">Active</option>
                        <option value="pending">Pending</option>
                        <option value="deactive">Deactive</option>
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
                        Avatar <span className="text-red-600"></span>
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
                      <Field
                        type="file"
                        name="image"
                        id="fileInput"
                        //id="image-preview"
                        onChange={(e: any) => handleChange(e)}
                      />
                      <img
                        id="image-preview"
                        src={image}
                        alt="preview"
                        className="mt-2 w-[50%] h-[50%]"
                        // id="avatar"
                      />
                      {/* <p className="text-green-400">size up to 5MB</p> */}

                      {submitCount ? (
                        errors.avatar ? (
                          <div className="text-danger mt-1">
                            {errors.avatar}
                          </div>
                        ) : null
                      ) : (
                        // (
                        //   <div className="text-success mt-1">It is fine!</div>
                        // )
                        ""
                      )}
                    </div>
                  </div>
                </div>

                <div className="flex justify-center py-6 mt-12">
                  {/* <button
                    type="reset"
                    className="btn btn-outline-dark rounded-xl px-8 mx-3 font-bold"
                  >
                    RESET
                  </button> */}

                  <button
                    type="submit"
                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828] mx-3"
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

export default EditUser;
