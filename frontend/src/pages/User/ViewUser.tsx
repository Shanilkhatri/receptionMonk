import { Link, NavLink, useNavigate } from "react-router-dom";
//import { DataTable, DataTableSortStatus } from "mantine-datatable";
import { useState, useEffect, useRef } from "react";
import sortBy from "lodash/sortBy";
import { useDispatch, useSelector } from "react-redux";
import { IRootState } from "../../store";
import {
  setPageTitle,
  setcurrentUserDataForUpdate,
} from "../../store/themeConfigSlice";
import Utility from "../../utility/utility";
import { HotTable } from "@handsontable/react";
import { registerAllModules } from "handsontable/registry";
import { textRenderer } from "handsontable/renderers/textRenderer";

import Handsontable from "handsontable";

import "handsontable/dist/handsontable.full.min.css";

registerAllModules();

// object of class utility
const utility = new Utility();
const appUrl = import.meta.env.VITE_APPURL;
const ViewUser = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  //const hotRef = useRef(null);
  const hotTableComponentRef = useRef<any>(null);

  useEffect(() => {
    if (hotTableComponentRef.current != null) {
      const handsontableInstance = hotTableComponentRef.current.hotInstance;
      const filterField = document.querySelector(
        "#filterField"
      ) as HTMLInputElement;

      filterField.addEventListener("keyup", function (event: any) {
        const filtersPlugin = handsontableInstance.getPlugin("filters");
        const columnSelector = document.getElementById(
          "columns"
        ) as HTMLInputElement;
        const columnValue = columnSelector.value;

        filtersPlugin.removeConditions(columnValue);
        filtersPlugin.addCondition(columnValue, "contains", [
          event.target.value,
        ]);
        filtersPlugin.filter();

        handsontableInstance.render();
      });
    }
  }, []);

  let searchFieldKeyupCallback: any;

  useEffect(() => {
    dispatch(setPageTitle("View User"));
  });

  const isDark =
    useSelector((state: IRootState) => state.themeConfig.theme) === "dark"
      ? true
      : false;
  const [items, setItems] = useState([]);

  const deleteRow = (id: any = null) => {
    if (window.confirm("Are you sure want to delete selected row ?")) {
      if (id) {
        setRecords(items.filter((user) => user.id !== id));
        setInitialRecords(items.filter((user) => user.id !== id));
        setItems(items.filter((user) => user.id !== id));
        setSearch("");
        setSelectedRecords([]);
      } else {
        let selectedRows = selectedRecords || [];
        const ids = selectedRows.map((d: any) => {
          return d.id;
        });
        const result = items.filter((d) => !ids.includes(d.id as never));
        setRecords(result);
        setInitialRecords(result);
        setItems(result);
        setSearch("");
        setSelectedRecords([]);

        setPage(1);
      }
    }
  };

  const [page, setPage] = useState(1);
  const PAGE_SIZES = [10, 20, 30, 50, 100];
  const [pageSize, setPageSize] = useState(PAGE_SIZES[0]);
  const [initialRecords, setInitialRecords] = useState(sortBy(items, "id"));
  const [records, setRecords] = useState(initialRecords);
  const [selectedRecords, setSelectedRecords] = useState<any>([]);

  const [search, setSearch] = useState("");
  const [sortStatus, setSortStatus] = useState<any>({
    columnAccessor: "userName",
    direction: "asc",
  });
  const fetchData = async () => {
    try {
      var token = utility.getCookieValue("exampleToken");
      var headers = new Headers();
      headers.append("Content-Type", "application/json");
      headers.append("Authorization", "bearer " + token);
      // Make API call or fetch data from wherever you need
      const response = await fetch(appUrl + "users", {
        method: "GET",
        headers: headers,
      });
      const data = await response.json();
      console.log(data);

      let arrayOfDesiredSet: any = [];
      data.Payload.forEach((item: any) => {
        let desiredDataSet = {
          id: 0,
          name: "",
          email: "",
          dob: "",
          accountType: "",
          status: "",
          avatar: "",
        };
        desiredDataSet["id"] = item.id;
        desiredDataSet["name"] = item.name;
        desiredDataSet["email"] = item.email;
        desiredDataSet["dob"] = item.dob;
        desiredDataSet["accountType"] = item.accountType;
        desiredDataSet["status"] = item.status;
        desiredDataSet["avatar"] = item.avatar;
        arrayOfDesiredSet.push(desiredDataSet);
      });
      // Update the state with the fetched data
      console.log("arrayOfDesiredSet", arrayOfDesiredSet);
      setItems(arrayOfDesiredSet);
      setInitialRecords(arrayOfDesiredSet);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };
  useEffect(() => {
    fetchData(); // Call the fetchData function when the component mounts
  }, []);

  useEffect(() => {
    setPage(1);
    /* eslint-disable react-hooks/exhaustive-deps */
  }, [pageSize]);

  useEffect(() => {
    const from: any = (page - 1) * pageSize;
    const to: any = from + pageSize;
    setRecords([...initialRecords.slice(from, to)]);
  }, [page, pageSize, initialRecords]);

  //   useEffect(() => {
  //     setInitialRecords(() => {
  //       return items.filter((item) => {
  //         return (
  //           // item.invoice.toLowerCase().includes(search.toLowerCase()) || //to be removed
  //           item.userName.toLowerCase().includes(search.toLowerCase()) ||
  //           item.userEmail.toLowerCase().includes(search.toLowerCase()) ||
  //           item.userDob.toLowerCase().includes(search.toLowerCase()) ||
  //           //item.amount.toLowerCase().includes(search.toLowerCase()) ||
  //           item.userAccType.toLowerCase().includes(search.toLowerCase()) ||
  //           // item.cname.toLowerCase().includes(search.toLowerCase()) || //to be removed
  //           item.userStatus.tooltip
  //             .toLowerCase()
  //             .includes(search.toLowerCase()) ||
  //           item.avatar
  //         );
  //       });
  //     });
  //   }, [search]);

  //   useEffect(() => {
  //     const container = hotTableComponentRef.current;

  //     //const hot = new Handsontable(container);

  //     return () => {
  //       hot.destroy();
  //     };
  //   }, []);

  const editButtonRenderer = (
    instance: any,
    td: any,
    row: any,
    col: any,
    prop: any,
    value: any,
    cellProperties: any
  ) => {
    // const rowId = instance.getDataAtRowProp(row, "id"); // Get the ID from the data
    const rowObject = instance.getDataAtRow(row);
    let desiredDataSet = {
      id: `${rowObject[0]}`,
      name: `${rowObject[1]}`,
      email: `${rowObject[2]}`,
      dob: `${rowObject[3]}`,
      status: `${rowObject[4]}`,
      accountType: `${rowObject[5]}`,
      avatar: `${rowObject[6]}`,
    };
    console.log("rowData: ", rowObject);
    td.innerHTML = `<button onclick='editRow( ${JSON.stringify(
      desiredDataSet
    )} )'>Edit</button>`;
    return td;
  };
  // Handle the edit action
  window.editRow = (rowData: any) => {
    // Access the row data and perform the desired action
    // calling dispatcher to set uodate user in state
    dispatch(setcurrentUserDataForUpdate(rowData));
    navigate("/edituser");

    // Add your logic here, such as opening a modal for editing
  };

  useEffect(() => {
    const data2 = sortBy(initialRecords, sortStatus.columnAccessor);
    setRecords(sortStatus.direction === "desc" ? data2.reverse() : data2);
    setPage(1);
  }, [sortStatus]);
  return (
    <div>
      <ul className="flex space-x-2 rtl:space-x-reverse">
        <li>
          <Link to="#" className="text-primary hover:underline">
            User
          </Link>
        </li>
        <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
          <span>View</span>
        </li>
      </ul>

      <div className="panel px-0 border-white-light dark:border-[#1b2e4b] py-5">
        <div className="text-xl font-bold text-dark dark:text-white text-center py-8">
          View User
        </div>

        <div className="invoice-table">
          <div className="datatables pagination-padding">
            {/* data table */}
            <div className="controlsQuickFilter flex mx-2">
              <label htmlFor="columns" className="selectColumn mr-4">
                Select a column:{" "}
              </label>
              <select name="columns" id="columns">
                <option value="0">userName</option>
                <option value="1">userid</option>
                <option value="2">userEmail</option>
                <option value="3">userDob</option>
                <option value="4">userAccType</option>
                <option value="5">avatar</option>
              </select>
            </div>
            <div className="controlsQuickFilter mx-2 my-5">
              <input id="filterField" type="text" placeholder="Filter" />
            </div>

            <HotTable
              ref={hotTableComponentRef}
              data={items}
              autoRowSize={true}
              columns={[
                {
                  title: "id",
                  type: "numeric",
                  data: "id",
                },
                {
                  title: "User Name",
                  type: "text",
                  data: "name",
                },
                {
                  title: "User Email",
                  type: "text",
                  data: "email",
                },
                {
                  title: "D.O.B",
                  type: "text",
                  data: "dob",
                },
                {
                  title: "Status",
                  type: "text",
                  data: "status",
                },
                {
                  title: "Account Type",
                  type: "text",
                  data: "accountType",
                },
                {
                  title: "Avatar",
                  data: "avatar",
                  renderer(
                    instance,
                    td,
                    row,
                    col,
                    prop,
                    value,
                    cellProperties
                  ) {
                    const img = document.createElement("img");
                    //apply condition to default image
                    // if(value == ""){
                    // img.src=
                    // }
                    img.src = import.meta.env.VITE_APPURL + value;
                    img.alt = "Not Available";
                    img.width = 90;
                    img.addEventListener("mousedown", (event) => {
                      event.preventDefault();
                    });

                    td.innerText = "";

                    td.appendChild(img);

                    return td;
                  },
                },
                {
                  title: "Action",
                  data: "edit",
                  renderer: editButtonRenderer,
                  readOnly: true,
                },
                {},
              ]}
              cw
              rowHeights={55}
              readOnly={true}
              colHeaders={true}
              // enable the column menu
              filters={true}
              className="exampleQuickFilter "
              licenseKey="non-commercial-and-evaluation" // for non-commercial use only
            />

            {/* <DataTable
              className={`${isDark} whitespace-nowrap table-hover`}
              records={records}
              columns={[
                {
                  accessor: "userName",
                  sortable: true,
                  render: ({ userName, id }) => (
                    <div className="flex items-center font-semibold">
                      <div className="p-0.5 bg-white-dark/30 rounded-full w-max ltr:mr-2 rtl:ml-2">
                        <img
                          className="h-8 w-8 rounded-full object-cover"
                          src={`/assets/images/profile-${id}.jpeg`}
                          alt=""
                        />
                      </div>
                      <div>{userName}</div>
                    </div>
                  ),
                },
                {
                  accessor: "userEmail",
                  sortable: true,
                },
                {
                  accessor: "userDob",
                  sortable: true,
                },
                {
                  accessor: "userAccType",
                  sortable: true,
                  render: ({ userAccType, id }) => (
                    <div className="font-semibold">{`${userAccType}`}</div>
                  ),
                },
                {
                  accessor: "userStatus",
                  sortable: true,
                  render: ({ userStatus }) => (
                    <span
                      className={`badge badge-outline-${userStatus.color} `}
                    >
                      {userStatus.tooltip}
                    </span>
                  ),
                },
                {
                  accessor: "action",
                  title: "Actions",
                  sortable: false,
                  textAlignment: "center",
                 </div> render: ({ id }) => (*/}
            <div className="flex gap-4 items-center w-max mx-auto">
              <NavLink to="/edituser" className="flex hover:text-info">
                <i className="bi bi-pen"></i>
                {/* <svg
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-4.5 h-4.5"
                >
                  <path
                    opacity="0.5"
                    d="M22 10.5V12C22 16.714 22 19.0711 20.5355 20.5355C19.0711 22 16.714 22 12 22C7.28595 22 4.92893 22 3.46447 20.5355C2 19.0711 2 16.714 2 12C2 7.28595 2 4.92893 3.46447 3.46447C4.92893 2 7.28595 2 12 2H13.5"
                    stroke="currentColor"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                  ></path>
                  <path
                    d="M17.3009 2.80624L16.652 3.45506L10.6872 9.41993C10.2832 9.82394 10.0812 10.0259 9.90743 10.2487C9.70249 10.5114 9.52679 10.7957 9.38344 11.0965C9.26191 11.3515 9.17157 11.6225 8.99089 12.1646L8.41242 13.9L8.03811 15.0229C7.9492 15.2897 8.01862 15.5837 8.21744 15.7826C8.41626 15.9814 8.71035 16.0508 8.97709 15.9619L10.1 15.5876L11.8354 15.0091C12.3775 14.8284 12.6485 14.7381 12.9035 14.6166C13.2043 14.4732 13.4886 14.2975 13.7513 14.0926C13.9741 13.9188 14.1761 13.7168 14.5801 13.3128L20.5449 7.34795L21.1938 6.69914C22.2687 5.62415 22.2687 3.88124 21.1938 2.80624C20.1188 1.73125 18.3759 1.73125 17.3009 2.80624Z"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  ></path>
                  <path
                    opacity="0.5"
                    d="M16.6522 3.45508C16.6522 3.45508 16.7333 4.83381 17.9499 6.05034C19.1664 7.26687 20.5451 7.34797 20.5451 7.34797M10.1002 15.5876L8.4126 13.9"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  ></path>
                </svg> */}
              </NavLink>
              {/* <NavLink to="" className="flex">  */}
              <button
                type="button"
                className="flex hover:text-danger"
                onClick={(e) => deleteRow(id)}
              >
                <svg
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5"
                >
                  <path
                    d="M20.5001 6H3.5"
                    stroke="currentColor"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                  ></path>
                  <path
                    d="M18.8334 8.5L18.3735 15.3991C18.1965 18.054 18.108 19.3815 17.243 20.1907C16.378 21 15.0476 21 12.3868 21H11.6134C8.9526 21 7.6222 21 6.75719 20.1907C5.89218 19.3815 5.80368 18.054 5.62669 15.3991L5.16675 8.5"
                    stroke="currentColor"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                  ></path>
                  <path
                    opacity="0.5"
                    d="M9.5 11L10 16"
                    stroke="currentColor"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                  ></path>
                  <path
                    opacity="0.5"
                    d="M14.5 11L14 16"
                    stroke="currentColor"
                    strokeWidth="1.5"
                    strokeLinecap="round"
                  ></path>
                  <path
                    opacity="0.5"
                    d="M6.5 6C6.55588 6 6.58382 6 6.60915 5.99936C7.43259 5.97849 8.15902 5.45491 8.43922 4.68032C8.44784 4.65649 8.45667 4.62999 8.47434 4.57697L8.57143 4.28571C8.65431 4.03708 8.69575 3.91276 8.75071 3.8072C8.97001 3.38607 9.37574 3.09364 9.84461 3.01877C9.96213 3 10.0932 3 10.3553 3H13.6447C13.9068 3 14.0379 3 14.1554 3.01877C14.6243 3.09364 15.03 3.38607 15.2493 3.8072C15.3043 3.91276 15.3457 4.03708 15.4286 4.28571L15.5257 4.57697C15.5433 4.62992 15.5522 4.65651 15.5608 4.68032C15.841 5.45491 16.5674 5.97849 17.3909 5.99936C17.4162 6 17.4441 6 17.5 6"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  ></path>
                </svg>
              </button>
              {/* </NavLink> */}
            </div>
            {/*
                  ),
                },
              ]}
              highlightOnHover
              totalRecords={initialRecords.length}
              recordsPerPage={pageSize}
              page={page}
              onPageChange={(p: any) => setPage(p)}
              recordsPerPageOptions={PAGE_SIZES}
              onRecordsPerPageChange={setPageSize}
              sortStatus={sortStatus}
              onSortStatusChange={setSortStatus}
              selectedRecords={selectedRecords}
              onSelectedRecordsChange={setSelectedRecords}
              paginationText={({ from, to, totalRecords }) =>
                `Showing  ${from} to ${to} of ${totalRecords} entries`
              }
            /> */}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ViewUser;
