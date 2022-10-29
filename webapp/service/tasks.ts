import {AxiosResponse} from "axios";
import {PaginatedResults, Task} from "../types/tasks";
import {getAxios} from "../utils/utils";

const axios = getAxios()
export const getTaskHistory = async (page = 0) : Promise<PaginatedResults<Task>> => {
    const tasks = await axios.get<{}, AxiosResponse< PaginatedResults<Task>>>('/api/tasks/history',{params:{page, size:25}})
    return tasks.data
}