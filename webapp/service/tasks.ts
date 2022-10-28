import axios, {AxiosResponse} from "axios";
import {PaginatedResults, Task} from "../types/tasks";

export const getTaskHistory = async (page = 0) : Promise<PaginatedResults<Task>> => {

    const tasks = await axios.get<{}, AxiosResponse< PaginatedResults<Task>>>('/api/tasks/history')
    return tasks.data
}