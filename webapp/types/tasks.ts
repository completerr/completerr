enum TaskStatus {
    Started = "started",
    Finished = "finished"
}
export interface PaginatedResults<T> {
    items: T[],
    page: number
    size: number
    max_page: number
    total_pages: number
    total: number
    last: boolean
    first: boolean
    visible: number
}
export type PaginationInfo = Omit<PaginatedResults<any>, "items">
export interface Task {
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    ID: number;
    type: string;
    status: string;
    started: Date;
    finished: Date;
}

export interface TvItem {
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    ID: number;
    name: string;
    available: boolean;
    title: string;
    last_searched: Date;
    sonarr_id: number;
    search_count: number;
    sonarr_series_id: number;
    series_title: string;
    season: number;
    season_episode_number: number;
    absolute_episode_number: number;
}

export interface Item {
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    ID: number;
    name: string;
    available: boolean;
    released: boolean;
    title: string;
    last_searched: string;
    tmdb_id: number;
    radarr_id: number;
    search_count: number;
}

export interface SearchRecord {
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    ID: number;
    task_id: number;
    task: Task;
    tv_item_id: number;
    tv_item: TvItem;
    item: Item;
    item_id: number;
}


export type RadarrSearchRecord = Omit<SearchRecord,"tv_item"| "tv_item_id">
export type SonarrSearchRecord = Omit<SearchRecord,"item"| "item_id">
export type ItemType = "radarr" | "sonarr"
