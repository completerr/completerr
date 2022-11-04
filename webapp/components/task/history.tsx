import React, {useEffect, useState} from 'react'
import {PaginatedResults, PaginationInfo, Task} from "../../types/tasks";
import {getTaskHistory} from "../../service/tasks";
import {Table} from "../table/table";
import {getRelativeDateString} from "../../utils/utils";

export const TaskHistory: React.FC<{}> = props => {
    const [taskHistory, setTaskHistory] = useState({} as PaginatedResults<Task>)
    const onPageClick = async (event: React.MouseEvent<HTMLButtonElement> | null, page: number) => {
        const res = await getTaskHistory(page)
        if (res !== taskHistory) {
            setTaskHistory(res)
        }
    }
    useEffect(() => {
        (async () => {
            const res = await getTaskHistory(0)
            if (res !== taskHistory) {
                setTaskHistory(res)
            }
        })()
    }, []);
    console.log(taskHistory)
    return taskHistory && taskHistory.items && (
        <Table headers={['Task Type', 'Status', 'Started', 'Finished']} pagination={taskHistory as PaginationInfo}
               data={taskHistory.items.map((item: Task) => [
                   item.type,
                   item.status,
                   getRelativeDateString(item.started),
                   item.status == "finished" ? getRelativeDateString(item.finished) : ""
               ])}
               paginationCallback={onPageClick} title={"Task History"}/>
    );
};

