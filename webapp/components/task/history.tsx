import React, {useEffect, useState} from 'react'
import {PaginatedResults, PaginationInfo, Task} from "../../types/tasks";
import {getTaskHistory} from "../../service/tasks";
import {OnPageClickParams, Table} from "../table/table";

export const TaskHistory: React.FC<{}> = props => {
    const [taskHistory, setTaskHistory] = useState({} as PaginatedResults<Task>)
    const onPageClick = async ({selected = 0}: OnPageClickParams) => {
        const res = await getTaskHistory(selected)
        if (res !== taskHistory) {
            setTaskHistory(res)
        }
    }
    useEffect(() => {
        onPageClick({selected: 0})
    }, []);

    return taskHistory && taskHistory.items && (
        <Table headers={['Task Type', 'Status', 'Started', 'Finished']} pagination={taskHistory as PaginationInfo}
               data={taskHistory.items.map((item: Task) => [item.type, item.status, item.started.toLocaleString(), item.finished.toLocaleString()])}
               paginationCallback={onPageClick} title={"Task History"}/>
    );
};

