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
import ReactDOM from "react-dom";
import App from "../../App";

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
    // if (hotTableComponentRef.current != null) {
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
          console.log("column value:", columnValue)
        filtersPlugin.removeConditions(columnValue);
        filtersPlugin.addCondition(columnValue, "contains", [
          event.target.value,
        ]);
        filtersPlugin.filter();
        handsontableInstance.render();
      });
    // }
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
          companyId: 0, // company id to be set in row object
        };
        desiredDataSet["id"] = item.id;
        desiredDataSet["name"] = item.name;
        desiredDataSet["email"] = item.email;
        desiredDataSet["dob"] = item.dob;
        desiredDataSet["accountType"] = item.accountType;
        desiredDataSet["status"] = item.status;
        desiredDataSet["avatar"] = item.avatar;
        desiredDataSet["companyId"] = item.companyId;
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
      companyId: `${rowObject[7]}`,
    };
   
    td.innerHTML = `<button onclick='editRow( ${JSON.stringify(
      desiredDataSet
    )} )'>Edit</button>`;
    return td;
  };
  // Handle the edit action
  window.editRow = (desiredDataSet: any) => {
    // Access the row data and perform the desired action
    // calling dispatcher to set uodate user in state
    dispatch(setcurrentUserDataForUpdate(desiredDataSet));
    navigate("/edituser");

    // Add your logic here, such as opening a modal for editing
  };

  const tableStyle = {
    // maxWidth: "800px", // Adjust the width as needed
    // margin: "0 auto", // Center the table horizontally
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
            {/* <div className="overflow-x-auto max-w-screen-lg mx-auto"> */}
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
                  title: "Company Id",
                  type: "numeric",
                  data: "companyId",
                  hidden: true, 
                },
                {
                  title: "Action",
                  data: "edit",
                  renderer: editButtonRenderer,
                  readOnly: true,
                },
              ]}
              height="auto"
              readOnly={true}
              colHeaders={true}
              stretchH="all"
              hiddenColumns={{
                columns: [7],
                indicators: true,
              }}
              //afterGetColHeader={(col, TH) => {}}
              // enable the column menu
              filters={true}
              className="exampleQuickFilter "
              licenseKey="non-commercial-and-evaluation" // for non-commercial use only
            />
            
          </div>
        </div>
      </div>
    </div>
  );
};
export default ViewUser;
