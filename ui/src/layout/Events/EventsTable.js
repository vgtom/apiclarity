import React, { useMemo } from 'react';
import { useHistory, useRouteMatch } from 'react-router-dom';
import Table, { utils } from 'components/Table';
import Tag from 'components/Tag';
import Icon, { ICON_NAMES } from 'components/Icon';
import StatusIndicator from 'components/StatusIndicator';
import { formatDate } from 'utils/utils';
import { API_TYPE_ITEMS } from 'layout/Inventory';

const EventsTable = ({filters, refreshTimestamp}) => {
    const columns = useMemo(() => [
        {
            Header: 'Time',
            id: "time",
            accessor: original => formatDate(original.time),
            width: 70
        },
        {
            Header: 'Method',
            id: "method",
            Cell: ({row}) => {
                const {method} = row.original;

                return (
                    <Tag>{method}</Tag>
                )
            },
            canSort: true,
            width: 40
        },
        {
            Header: 'Path',
            id: "path",
            accessor: "path"
        },
        {
            Header: 'Status Code',
            id: "statusCode",
            Cell: ({row}) => {
                const {statusCode} = row.original;

                return (
                    <StatusIndicator title={statusCode} isError={statusCode >= 400} />
                )
            },
            canSort: true,
            width: 40
        },
        {
            Header: 'Source',
            id: "sourceIP",
            accessor: "sourceIP",
            width: 50
        },
        {
            Header: 'Destination',
            id: "destinationIP",
            accessor: "destinationIP",
            width: 50
        },
        {
            Header: 'Destination Port',
            id: "destinationPort",
            accessor: "destinationPort",
            width: 55
        },
        {
            Header: 'Spec Diff',
            id: "hasSpecDiff",
            Cell: ({row}) => {
                const {hasProvidedSpecDiff, hasReconstructedSpecDiff} = row.original;
                
                return hasProvidedSpecDiff || hasReconstructedSpecDiff ?
                    <Icon name={ICON_NAMES.ALERT} className="specs-diff-alert-icon" /> : <utils.EmptyValue />;
            },
            canSort: true,
            width: 40
        },
        {
            Header: 'Host',
            id: "hostSpecName",
            accessor: "hostSpecName"
        },
        {
            Header: 'Type',
            id: "apiType",
            accessor: original => {
                const typeItem = API_TYPE_ITEMS[original.apiType];

                return !!typeItem ? typeItem.label : null;
            },
            width: 30
        }
    ], []);

    const history = useHistory();
    const {path} = useRouteMatch();
    
    return (
        <Table
            columns={columns}
            paginationItemsName="APIs"
            url="apiEvents"
            defaultSortBy={[{id: "time", desc: true}]}
            filters={filters}
            onLineClick={({id}) => history.push(`${path}/${id}`)}
            noResultsTitle="API events"
            refreshTimestamp={refreshTimestamp}
        />
    )
}

export default EventsTable;