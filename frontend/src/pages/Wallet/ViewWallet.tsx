import { Link, NavLink } from 'react-router-dom';
import { DataTable, DataTableSortStatus } from 'mantine-datatable';
import { useState, useEffect } from 'react';
import sortBy from 'lodash/sortBy';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';

const ViewWallet = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('View Wallet'));
    });

    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;
    const [items, setItems] = useState([
        {
            id: 1,
            charge: 'Plan 199',
            amount: '₹ 199',
            date: '15 Dec 2020',
            cname: 'HP',
            reason: 'IVR Sales',
        },
        {
            id: 2,
            charge: 'Plan 99',
            amount: '₹ 99',
            date: '10 Feb 2022',
            cname: 'Dell',
            reason: 'IVR Sales',
        },
        {
            id: 3,
            charge: 'Plan 299',
            amount: '₹ 299',
            date: '17 Mar 2023',
            cname: 'Asus',
            reason: 'IVR Billing',
        },
        {
            id: 4,
            charge: 'Plan 249',
            amount: '₹ 249',
            date: '15 Dec 2020',
            cname: 'Dell',
            reason: 'IVR Support',
        },
        {
            id: 5,
            charge: 'Plan 165',
            amount: '₹ 165',
            date: '10 Feb 2022',
            cname: 'Asus',
            reason: 'IVR Support',
        },
        {
            id: 6,
            charge: 'Plan 399',
            amount: '₹ 399',
            date: '17 Mar 2023',
            cname: 'Lenovo',
            reason: 'IVR Billing',
        },
        {
            id: 7,
            charge: 'Plan 99',
            amount: '₹ 99',
            date: '15 Dec 2020',
            cname: 'HP',
            reason: 'IVR Sales',
        },
        {
            id: 8,
            charge: 'Plan 899',
            amount: '₹ 899',
            date: '17 Mar 2023',
            cname: 'Dell',
            reason: 'IVR Support',
        },
        {
            id: 9,
            charge: 'Plan 99',
            amount: '₹ 99',
            date: '15 Dec 2020',
            cname: 'Asus',
            reason: 'IVR Billing',
        },
        {
            id: 10,
            charge: 'Plan 1299',
            amount: '₹ 1299',
            date: '10 Feb 2022',
            cname: 'Lenovo',
            reason: 'IVR Sales',
        },
        {
            id: 11,
            charge: 'Plan 149',
            amount: '₹ 149',
            date: '06 July 2018',
            cname: 'HP',
            reason: 'IVR Support',
        },
        {
            id: 12,
            charge: 'Plan 99',
            amount: '₹ 99',
            date: '15 June 2020',
            cname: 'Dell',
            reason: 'IVR Billing',
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
        columnAccessor: 'firstName',
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
                    item.charge.toLowerCase().includes(search.toLowerCase()) ||
                    item.amount.toLowerCase().includes(search.toLowerCase()) ||
                    item.date.toLowerCase().includes(search.toLowerCase()) ||
                    item.cname.toLowerCase().includes(search.toLowerCase()) ||
                    item.reason.toLowerCase().includes(search.toLowerCase())
                );
            });
        });
    }, [search]);

    useEffect(() => {
        const data2 = sortBy(initialRecords, sortStatus.columnAccessor);
        setRecords(sortStatus.direction === 'desc' ? data2.reverse() : data2);
        setPage(1);
    }, [sortStatus]);   
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Wallet
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>View</span>
                </li>
            </ul>

            <div className="panel px-0 border-white-light dark:border-[#1b2e4b] py-5">
                
                    <div className='text-xl font-bold text-dark dark:text-white text-center py-8'>View Wallet</div>

                    <div className="invoice-table">
                        <div className="datatables pagination-padding">
                            <DataTable
                                className={`${isDark} whitespace-nowrap table-hover`}
                                records={records}
                                columns={[
                                    
                                    
                                    {
                                        accessor: 'charge',
                                        sortable: true,
                                    },
                                    {
                                        accessor: 'amount',
                                        sortable: true,
                                    },
                                    {
                                        accessor: 'date',
                                        sortable: true,
                                        render: ({ date, id }) => <div className="font-semibold">{`${date}`}</div>,
                                    },
                                    {
                                        accessor: 'company name',
                                        sortable: true,
                                        render: ({ cname, id }) => <div className="font-semibold">{`${cname}`}</div>,
                                    },
                                    {
                                        accessor: 'reason',
                                        sortable: true,
                                        render: ({ reason, id }) => <div className="font-semibold">{`${reason}`}</div>,
                                    },
                                    {
                                        accessor: 'action',
                                        title: 'Actions',
                                        sortable: false,
                                        textAlignment: 'center',
                                        render: ({ id }) => (
                                            <div className="flex gap-4 items-center w-max mx-auto">
                                                <NavLink to="/editwallet" className="flex hover:text-info">
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

export default ViewWallet;