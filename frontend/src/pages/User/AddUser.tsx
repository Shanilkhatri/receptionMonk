import { Link, useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { IRootState } from "../../store";
import { setPageTitle } from "../../store/themeConfigSlice";
import * as Yup from "yup";
import { useEffect, useState } from "react";
import { Field, Form, Formik } from "formik";
import Swal from "sweetalert2";
import store from "../../store";
import Utility from "../../utility/utility";
// object of class utility
const utility = new Utility();
const appUrl = import.meta.env.VITE_APPURL;

const AddUser = () => {
  const dispatch = useDispatch();
  const [image, setImage] = useState<string | undefined>();

  useEffect(() => {
    dispatch(setPageTitle("Add User"));
  });

  const navigate = useNavigate();

  const submitForm = (e: any) => {
    // console.log("userData");
    const toast = Swal.mixin({
      toast: true,
      position: "top-end",
      showConfirmButton: false,
      timer: 3000,
    });
    toast.fire({
      icon: "success",
      title: "User Added Successfully",
      padding: "10px 20px",
    });
    // navigate("/viewuser");
  };

  async function addUser(data: any) {
    console.log("adding user");

    let img = await uploadImageAndReturnUrl("fileInput", "kyc");
    console.log("images", img);
    const userData = {
      name: data.userName,
      email: data.userEmail,
      dob: data.userDob,
      status: data.userStatus,
      accountType: data.userAccType,
      avatar: img,
      companyId: store.getState().themeConfig.hydrateCookie.companyId,
      iswizardcomplete: "kyc",
    };

    console.log(userData);
    const ok = await utility.sendRequestPutOrPost(userData, "users", "PUT");
    if (ok) {
      //do what you want if successfully added the data
      console.log("SuccessFully added");
      // check Token & update cookies
      // calling_token_check();
    } else {
      navigate("/auth/SignIn");
    }
  }
  function validateImgBeforeUpload(imageFile: File): boolean {
    var allowedExtensions = ["jpg", "jpeg", "png"]; // Allowed image extensions
    var maxFileSize = 5 * 1024 * 1024; // Maximum file size in bytes (5MB)
    if (!imageFile.name) {
      console.error("File name is undefined");
      return false;
    }
    // Check the file extension
    const fileParts = imageFile.name.split(".");
    if (fileParts.length === 1) {
      console.error("File name does not have an extension");
      return false;
    }
    // Check the file extension
    var fileExtension = fileParts.pop()!.toLowerCase();
    if (!allowedExtensions.includes(fileExtension)) {
      return false;
    }
    // Check the file size
    if (imageFile.size > maxFileSize) {
      console.log("failure", "Image size should be under 5mb");
      return false;
    }
    return true;
  }
  async function uploadImageAndReturnUrl(
    imageElement: string,
    modulename: string
  ): Promise<string> {
    const imageUploadURL = appUrl + "kycfileupload";

    const imageInput = document.getElementById(
      imageElement
    ) as HTMLInputElement | null;
    let imageFile: File | undefined;

    if (
      imageInput !== null &&
      imageInput.files &&
      imageInput.files.length > 0
    ) {
      imageFile = imageInput.files[0];
    }
    const formData = new FormData();
    if (imageFile != undefined && imageInput != null) {
      const valid = validateImgBeforeUpload(imageFile);
      if (!valid) {
        console.error("not a valid image is uploaded:");
        return "";
      }
      formData.append("image", imageFile || "");
      formData.append("modulename", modulename);
    }
    try {
      var response = await fetch(imageUploadURL, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${utility.getCookieValue("exampleToken")}`,
        },
        body: formData,
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data = await response.json();
      return data.Payload;
    } catch (error) {
      console.error("Error during image upload:", error);
      return ""; // or throw an error, depending on your logic
    }
  }

  const isRtl =
    useSelector((state: IRootState) => state.themeConfig.rtlClass) === "rtl"
      ? true
      : false;

  function handleChange(e: any) {
    // console.log(e.target.files);
    setImage(URL.createObjectURL(e.target.files[0]));
    //setImage(e.target.files[0]);
  }

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

      <Formik
        initialValues={{
          userName: "",
          userEmail: "",
          userDob: "",
          userAccType: "",
          userStatus: "",
          avatar: "",
        }}
        validationSchema={Yup.object().shape({
          userName: Yup.string().required("Please fill User Name"),
          userEmail: Yup.string()
            .email("Invalid Email Address")
            .required("Please fill Email"),
          userDob: Yup.string().required("Please fill User DOB"),
          userAccType: Yup.string().required("Please select User Type"),
          userStatus: Yup.string().required("Please select User Status"),
          avatar: Yup.string(), //mixed(),
        })}
        onSubmit={(values, { setSubmitting }) => {
          setTimeout(() => {
            console.log(values);
            addUser(values); // to see the data
            setSubmitting(false);
          }, 400);
        }}
      >
        {({ errors, submitCount, touched }) => (
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
                        <option value="select">Select</option>
                        <option value="active">Active</option>
                        <option value="pending">Pending</option>
                        <option value="deactivate">Deactive</option>
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
                      <input
                        type="file"
                        id="fileInput"
                        onChange={handleChange}
                      />
                      <img
                        src={image}
                        alt=""
                        className="mt-2 w-[50%] h-[50%]"
                      />
                      <p className="text-green-400">size up to 5MB,</p>

                      {submitCount ? (
                        errors.avatar ? (
                          <div className="text-danger mt-1">
                            {errors.avatar}
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
                    onClick={(e) => {
                      if (touched.userName && !errors.userName) {
                        submitForm(e);
                      } else if (touched.userEmail && !errors.userEmail) {
                        submitForm(e);
                      } else if (touched.userDob && !errors.userDob) {
                        submitForm(e);
                      } else if (touched.userAccType && !errors.userAccType) {
                        submitForm(e);
                      } else if (touched.avatar && !errors.avatar) {
                        submitForm(e);
                      } else if (touched.userStatus && !errors.userStatus) {
                        submitForm(e);
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

export default AddUser;
