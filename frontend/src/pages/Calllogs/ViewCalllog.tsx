import { Link, NavLink } from 'react-router-dom';
import { DataTable, DataTableSortStatus } from 'mantine-datatable';
import { useState, useEffect } from 'react';
import sortBy from 'lodash/sortBy';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';

const ViewCalllog = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('View Call Logs'));
    });

    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;
    const [items, setItems] = useState([
        {
            id: 1,
            callfrom: 'John',
            callto: 'ABC Hospital',
            callloc: 'Ranjhi',
            calltime: '12:08:25',
            callext: 'No',
        },
        {
            id: 2,
            callfrom: 'Methew',
            callto: 'XYZ Operation',
            callloc: 'Madan Mahal',
            calltime: '02:02:55',
            callext: 'Yes',
        },
        {
            id: 3,
            callfrom: 'Victoria',
            callto: 'PQR Sales',
            callloc: 'Khamaria',
            calltime: '22:44:05',
            callext: 'No',
        },
        {
            id: 4,
            callfrom: 'Lily',
            callto: 'XYZ Operation',
            callloc: 'Civil Lines',
            calltime: '20:05:52',
            callext: 'Yes',
        },
        {
            id: 5,
            callfrom: 'Mosque',
            callto: 'PQR Sales',
            callloc: 'Vijay Nagar',
            calltime: '15:08:33',
            callext: 'No',
        },
        {
            id: 6,
            callfrom: 'Dyna',
            callto: 'ABC Hospital',
            callloc: 'Sarafa',
            calltime: '07:07:15',
            callext: 'Yes',
        },
        {
            id: 7,
            callfrom: 'Mortin',
            callto: 'XYZ Operation',
            callloc: 'Adarsh Nagar',
            calltime: '07:22:05',
            callext: 'No',
        },
        {
            id: 8,
            callfrom: 'Henry',
            callto: 'PQR Sales',
            callloc: 'Adhartal',
            calltime: '12:12:25',
            callext: 'No',
        },
        {
            id: 9,
            callfrom: 'Shally',
            callto: 'ABC Hospital',
            callloc: 'Karmeta',
            calltime: '13:00:00',
            callext: 'Yes',
        },
        {
            id: 10,
            callfrom: 'Headen',
            callto: 'XYZ Operation',
            callloc: 'Bilhari',
            calltime: '01:00:05',
            callext: 'No',
        },
        {
            id: 11,
            callfrom: 'John',
            callto: 'PQR Sales',
            callloc: 'Sadar',
            calltime: '02:02:05',
            callext: 'No',
        },
        {
            id: 12,
            callfrom: 'John',
            callto: 'ABC Hospital',
            callloc: 'Napier Town',
            calltime: '05:52:05',
            callext: 'Yes',
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
                    item.callfrom.toLowerCase().includes(search.toLowerCase()) ||
                    item.callto.toLowerCase().includes(search.toLowerCase()) ||
                    item.callloc.toLowerCase().includes(search.toLowerCase()) ||
                    item.calltime.toLowerCase().includes(search.toLowerCase()) ||
                    item.callext.toLowerCase().includes(search.toLowerCase())
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
                        Call Logs
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>View</span>
                </li>
            </ul>

            <div className="panel px-0 border-white-light dark:border-[#1b2e4b] py-5">                
                <div className='text-xl font-bold text-dark dark:text-white text-center py-8'>View Call Logs</div>

                <div className="invoice-table">
                    <div className="datatables pagination-padding">
                        <DataTable
                            className={`${isDark} whitespace-nowrap table-hover`}
                            records={records}
                            columns={[
                                
                                
                                {
                                    accessor: 'Call From',
                                    sortable: true,
                                    render: ({ callfrom, id }) => <div className="font-semibold">{`${callfrom}`}</div>,
                                },
                                {
                                    accessor: 'Call To',
                                    sortable: true,
                                    render: ({ callto, id }) => <div className="font-semibold">{`${callto}`}</div>,
                                },
                                {
                                    accessor: 'Call Place At',
                                    sortable: true,
                                    render: ({ callloc, id }) => <div className="font-semibold">{`${callloc}`}</div>,
                                },
                                {
                                    accessor: 'Call Duration',
                                    sortable: true,
                                    render: ({ calltime, id }) => <div className="font-semibold">{`${calltime}`}</div>,
                                },
                                {
                                    accessor: 'Call Extension',
                                    sortable: true,
                                    render: ({ callext, id }) => <div className="font-semibold">{`${callext}`}</div>,
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

export default ViewCalllog;