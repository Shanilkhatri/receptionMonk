import { Link, NavLink } from 'react-router-dom';
import { DataTable, DataTableSortStatus } from 'mantine-datatable';
import { useState, useEffect } from 'react';
import sortBy from 'lodash/sortBy';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';

const OrderDetails = () => {
const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Order Details'));
    });

    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;
    const [items, setItems] = useState([
        {
            ordid: '#1234234',
            phnno: '+91 74125 89632',            
            sipsname: 'Server Name 1',
            sipuname: 'K12002',
            sippwd: '145236',
            sipport: 'Port 8080',
            ivrenabled: 'YES',
            ivrflow: 'Operation',
            maxalloweduser: "2 lac",
            maxalloweddept: "50",
        },
        {
            ordid: '#34534532',
            phnno: '+91 52147 36528',
            sipsname: 'Server Name 2',
            sipuname: 'J45782',           
            sippwd: '232323',
            sipport: 'Port 8080',
            ivrenabled: 'NO',
            ivrflow: 'Support',
            maxalloweduser: "1 lac",
            maxalloweddept: "20",
        },  
        {
            ordid: '#54454334',
            phnno: '+91 74125 89632',
            sipsname: 'Server Name 1',
            sipuname: 'K12002',
            sippwd: '145236',
            sipport: 'Port 8080',
            ivrenabled: 'YES',
            ivrflow: 'Sales',
            maxalloweduser: "5 lac",
            maxalloweddept: "200",
        },
        {
            ordid: '#34553534',
            phnno: '+91 52147 36528',           
            sipsname: 'Server Name 2',
            sipuname: 'J45782',           
            sippwd: '232323',
            sipport: 'Port 8080',
            ivrenabled: 'YES',
            ivrflow: 'Support',
            maxalloweduser: "1 lac",
            maxalloweddept: "20",
        },  
        {
            ordid: '#86786786',
            phnno: '+91 74125 89632',
            sipsname: 'Server Name 1',
            sipuname: 'K12002',
            sippwd: '145236',
            sipport: 'Port 8080',
            ivrenabled: 'NO',
            ivrflow: 'Sales',
            maxalloweduser: "5 lac",
            maxalloweddept: "100",
        },
        {
            ordid: '#8978978',
            phnno: '+91 52147 36528',
            sipsname: 'Server Name 2',
            sipuname: 'J45782',           
            sippwd: '232323',
            sipport: 'Port 8080',
            ivrenabled: 'YES',
            ivrflow: 'Support',
            maxalloweduser: "1 lac",
            maxalloweddept: "20",
        },  
        {
            ordid: '#0967867867',
            phnno: '+91 74125 89632',
            sipsname: 'Server Name 1',
            sipuname: 'K12002',
            sippwd: '145236',
            sipport: 'Port 8080',
            ivrenabled: 'NO',
            ivrflow: 'Support',
            maxalloweduser: "1 lac",
            maxalloweddept: "20",
        },
        {
            ordid:'#7656756756',
            phnno: '+91 52147 36528',
            sipsname: 'Server Name 2',
            sipuname: 'J45782',           
            sippwd: '232323',
            sipport: 'Port 8080',
            ivrenabled: 'YES',
            ivrflow: 'Sales',
            maxalloweduser: "5 lac",
            maxalloweddept: "100",
        },       
        {
            ordid: '#6756757',
            phnno: '+91 74125 89632',
            sipsname: 'Server Name 1',
            sipuname: 'K12002',
            sippwd: '145236',
            sipport: 'Port 8080',
            ivrenabled: 'YES',
            ivrflow: 'Support',
            maxalloweduser: "1 lac",
            maxalloweddept: "20",
        },
        {
            ordid: '#78678678',
            phnno: '+91 52147 36528',
            sipsname: 'Server Name 2',
            sipuname: 'J45782',           
            sippwd: '232323',
            sipport: 'Port 8080',
            ivrenabled: 'NO',
            ivrflow: 'Support',
            maxalloweduser: "1 lac",
            maxalloweddept: "20",
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
                    item.phnno.toLowerCase().includes(search.toLowerCase()) ||
                    item.sipsname.toLowerCase().includes(search.toLowerCase()) ||
                    item.sipuname.toLowerCase().includes(search.toLowerCase()) ||
                    item.sippwd.toLowerCase().includes(search.toLowerCase()) ||
                    item.sipport.toLowerCase().includes(search.toLowerCase()) ||
                    item.ivrenabled.toLowerCase().includes(search.toLowerCase()) ||
                    item.ivrflow.toLowerCase().includes(search.toLowerCase()) ||
                    item.maxalloweduser.toLowerCase().includes(search.toLowerCase()) ||
                    item.maxalloweddept.toLowerCase().includes(search.toLowerCase()) 
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
                        Extension
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>View</span>
                </li>
            </ul>

            <div className="panel px-0 border-white-light dark:border-[#1b2e4b] py-5">
                
                    <div className='text-xl font-bold text-dark dark:text-white text-center py-8'>Order Details</div>

                    <div className="invoice-table">
                        <div className="datatables pagination-padding">
                            <DataTable
                                className={`${isDark} whitespace-nowrap table-hover`}
                                records={records}
                                columns={[
                                    
                                    {
                                        accessor: 'Order ID',
                                        sortable: true,
                                        render: ({ ordid, id }) => (
                                            <div className="flex items-center font-semibold">                                                
                                                <div>{ordid}</div>
                                            </div>
                                        ),
                                    },
                                    {
                                        accessor: 'Phone Number',
                                        sortable: true,
                                        render: ({ phnno, id }) => <div className="font-semibold">{`${phnno}`}</div>,
                                    },                                    
                                    {
                                        accessor: 'SIP Server Name',
                                        sortable: true,
                                        titleClassName: 'text-right',
                                        render: ({ sipsname, id }) => <div className="font-semibold">{`${sipsname}`}</div>,
                                    },
                                    {
                                        accessor: 'Sip Username',
                                        sortable: true,
                                        render: ({ sipuname, id }) => <div className="font-semibold">{`${sipuname}`}</div>,
                                    },
                                    {
                                        accessor: 'Sip Password',
                                        sortable: true,
                                        render: ({ sippwd, id }) => <div className="font-semibold">{`${sippwd}`}</div>,
                                    },
                                    {
                                        accessor: 'Sip Port',
                                        sortable: true,
                                        render: ({ sipport, id }) => <div className="font-semibold">{`${sipport}`}</div>,
                                    },
                                    {
                                        accessor: 'Is IVR Enabled',
                                        sortable: true,
                                        render: ({ ivrenabled, id }) => <div className="font-semibold">{`${ivrenabled}`}</div>,
                                    },
                                    {
                                        accessor: 'IVR Flow',
                                        sortable: true,
                                        render: ({ ivrflow, id }) => <div className="font-semibold">{`${ivrflow}`}</div>,
                                    },
                                    {
                                        accessor: 'Max User Allowed',
                                        sortable: true,
                                        render: ({ maxalloweduser, id }) => <div className="font-semibold">{`${maxalloweduser}`}</div>,
                                    },
                                    {
                                        accessor: 'Max Dept. Allowed',
                                        sortable: true,
                                        render: ({ maxalloweddept, id }) => <div className="font-semibold">{`${maxalloweddept}`}</div>,
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

export default OrderDetails;