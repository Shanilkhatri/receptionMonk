import { Link, NavLink } from 'react-router-dom';
import { DataTable, DataTableSortStatus } from 'mantine-datatable';
import { useState, useEffect } from 'react';
import sortBy from 'lodash/sortBy';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';

const ViewOrder = () => {
const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('View Order'));
    });

    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;
    const [items, setItems] = useState([
        {
            id: 1,
            ordfrom: '14 Dec 2022',
            ordto: '23 Jan 2023',
            price: '₹ 1245',
            buyer: 'Intel',
            status: { tooltip: 'Paid', color: 'success' },
        },
        {
            id: 2,
            ordfrom: '06 May 2019',
            ordto: '02 June 2019',
            price: '₹ 5623',
            buyer: 'HP',
            status: { tooltip: 'Unpaid', color: 'danger' },     
        },  
        {
            id: 3,
            ordfrom: '18 Oct 2022',
            ordto: '20 Nov 2022',
            price: '₹ 5896',
            buyer: 'Asus',
            status: { tooltip: 'Unpaid', color: 'danger' },  
        },
        {
            id: 4,
            ordfrom: '12 Feb 2018',
            ordto: '06 Apr 2018',
            price: '₹ 7458',
            buyer: 'Zebronics',
            status: { tooltip: 'Unpaid', color: 'danger' },   
        },  
        {
            id: 5,
            ordfrom: '18 Oct 2022',
            ordto: '20 Nov 2022',
            price: '₹ 1452',
            buyer: 'Dell',
            status: { tooltip: 'Paid', color: 'success' },
        },
        {
            id: 6,
            ordfrom: '12 Feb 2018',
            ordto: '06 Apr 2018',
            price: '₹ 2563',
            buyer: 'Lenovo',
            status: { tooltip: 'Paid', color: 'success' },   
        },  
        {
            id: 7,
            ordfrom: '12 Feb 2018',
            ordto: '06 Apr 2018',
            price: '₹ 8574',
            buyer: 'HP',
            status: { tooltip: 'Unpaid', color: 'danger' },  
        },
        {
            id: 8,
            ordfrom: '18 Oct 2022',
            ordto: '20 Nov 2022',
            price: '₹ 1452',
            buyer: 'Asus',
            status: { tooltip: 'Unpaid', color: 'danger' },      
        },       
        {
            id: 9,
            ordfrom: '06 May 2019',
            ordto: '02 June 2019',
            price: '₹ 2563',
            buyer: 'Dell',
            status: { tooltip: 'Paid', color: 'success' },
        },
        {
            id: 10,
            ordfrom: '14 Dec 2022',
            ordto: '23 Jan 2023',
            price: '₹ 8569',
            buyer: 'Lenovo',
            status: { tooltip: 'Unpaid', color: 'danger' },     
        },        
    ]);

    const deleteRow = (id: any = null) => {
        if (window.confirm('Are you sure want to delete selected row ?')) {
            if (id) {
                setRecords(items.filter((user) => user.id !== id));
                setInitialRecords(items.filter((user) => user.id !== id));
                setItems(items.filter((user) => user.id !== id));
                setSearch('');
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
                setSearch('');
                setSelectedRecords([]);
                setPage(1);
            }
        }
    };

    const [page, setPage] = useState(1);
    const PAGE_SIZES = [10, 20, 30, 50, 100];
    const [pageSize, setPageSize] = useState(PAGE_SIZES[0]);
    const [initialRecords, setInitialRecords] = useState(sortBy(items, 'invoice'));
    const [records, setRecords] = useState(initialRecords);
    const [selectedRecords, setSelectedRecords] = useState<any>([]);

    const [search, setSearch] = useState('');
    const [sortStatus, setSortStatus] = useState<DataTableSortStatus>({
        columnAccessor: 'extname',
        direction: 'asc',
    });

    useEffect(() => {
        setPage(1);
        /* eslint-disable react-hooks/exhaustive-deps */
    }, [pageSize]);

    useEffect(() => {
        const from = (page - 1) * pageSize;
        const to = from + pageSize;
        setRecords([...initialRecords.slice(from, to)]);
    }, [page, pageSize, initialRecords]);

    useEffect(() => {
        setInitialRecords(() => {
            return items.filter((item) => {
                return (                    
                    item.ordfrom.toLowerCase().includes(search.toLowerCase()) ||
                    item.ordto.toLowerCase().includes(search.toLowerCase()) ||
                    item.price.toLowerCase().includes(search.toLowerCase()) ||
                    item.buyer.toLowerCase().includes(search.toLowerCase()) ||
                    item.status.tooltip.toLowerCase().includes(search.toLowerCase())
                );
            });
        });
    }, [search]);

    // useEffect(() => {
    //     const data2 = sortBy(initialRecords, sortStatus.columnAccessor);
    //     setRecords(sortStatus.direction === 'desc' ? data2.reverse() : data2);
    //     setPage(1);
    // }, [sortStatus]);   
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Order
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>View</span>
                </li>
            </ul>

            <div className="panel px-0 border-white-light dark:border-[#1b2e4b] py-5">
                
                    <div className='text-xl font-bold text-dark dark:text-white text-center py-8'>View Order</div>

                    <div className="invoice-table">
                        <div className="datatables pagination-padding">
                            <DataTable
                                className={`${isDark} whitespace-nowrap table-hover`}
                                records={records}
                                columns={[
                                    
                                    {
                                        accessor: 'Order From',
                                        sortable: true,
                                        render: ({ ordfrom, id }) => (
                                            <div className="flex items-center font-semibold">                                                
                                                <div>{ordfrom}</div>
                                            </div>
                                        ),
                                    },
                                    {
                                        accessor: 'Order Upto',
                                        sortable: true,
                                        render: ({ ordto, id }) => <div className="font-semibold">{`${ordto}`}</div>,
                                    },
                                    {
                                        accessor: 'Order Price',
                                        sortable: true,
                                        render: ({ price, id }) => <div className="font-semibold">{`${price}`}</div>,
                                    },
                                    {
                                        accessor: 'Buyer',
                                        sortable: true,
                                        titleClassName: 'text-right',
                                        render: ({ buyer, id }) => <div className="font-semibold">{`${buyer}`}</div>,
                                    },
                                    {
                                        accessor: 'Status',
                                        sortable: true,
                                        render: ({ status }) => <span className={`badge badge-outline-${status.color} `}>{status.tooltip}</span>,
                                    },
                                    {
                                        accessor: 'action',
                                        title: 'Actions',
                                        sortable: false,
                                        textAlignment: 'center',
                                        render: ({ id }) => (
                                            <div className="flex gap-4 items-center w-max mx-auto">
                                                <NavLink to="/editorder" className="flex hover:text-info">
                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="w-4.5 h-4.5">
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
                                                    </svg>
                                                </NavLink>     

                                                <NavLink to="/orderdetails" className="flex hover:text-primary">
                                                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                        <path opacity="0.5" d="M3.27489 15.2957C2.42496 14.1915 2 13.6394 2 12C2 10.3606 2.42496 9.80853 3.27489 8.70433C4.97196 6.49956 7.81811 4 12 4C16.1819 4 19.028 6.49956 20.7251 8.70433C21.575 9.80853 22 10.3606 22 12C22 13.6394 21.575 14.1915 20.7251 15.2957C19.028 17.5004 16.1819 20 12 20C7.81811 20 4.97196 17.5004 3.27489 15.2957Z" stroke="currentColor" stroke-width="1.5"></path><path d="M15 12C15 13.6569 13.6569 15 12 15C10.3431 15 9 13.6569 9 12C9 10.3431 10.3431 9 12 9C13.6569 9 15 10.3431 15 12Z" stroke="currentColor" stroke-width="1.5"></path>
                                                    </svg>
                                                </NavLink>

                                                {/* <NavLink to="" className="flex"> */}
                                                <button type="button" className="flex hover:text-danger" onClick={(e) => deleteRow(id)}>
                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="h-5 w-5">
                                                        <path d="M20.5001 6H3.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"></path>
                                                        <path
                                                            d="M18.8334 8.5L18.3735 15.3991C18.1965 18.054 18.108 19.3815 17.243 20.1907C16.378 21 15.0476 21 12.3868 21H11.6134C8.9526 21 7.6222 21 6.75719 20.1907C5.89218 19.3815 5.80368 18.054 5.62669 15.3991L5.16675 8.5"
                                                            stroke="currentColor"
                                                            strokeWidth="1.5"
                                                            strokeLinecap="round"
                                                        ></path>
                                                        <path opacity="0.5" d="M9.5 11L10 16" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"></path>
                                                        <path opacity="0.5" d="M14.5 11L14 16" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"></path>
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
                                        ),
                                    },
                                ]}
                                highlightOnHover
                                totalRecords={initialRecords.length}
                                recordsPerPage={pageSize}
                                page={page}
                                onPageChange={(p) => setPage(p)}
                                recordsPerPageOptions={PAGE_SIZES}
                                onRecordsPerPageChange={setPageSize}
                                sortStatus={sortStatus}
                                onSortStatusChange={setSortStatus}
                                selectedRecords={selectedRecords}
                                onSelectedRecordsChange={setSelectedRecords}
                                paginationText={({ from, to, totalRecords }) => `Showing  ${from} to ${to} of ${totalRecords} entries`}
                            />
                        </div>
                    </div>
                    
            </div>
        </div>
    );
};

export default ViewOrder;