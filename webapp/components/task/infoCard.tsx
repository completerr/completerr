import {Card, Fab} from "@mui/material";
import Typography from "@mui/material/Typography";
import {getRelativeDateString} from "../../utils/utils";
import Grid from "@mui/material/Grid";

export interface TaskInfoCardProps {
    name: string,
    prevRun: Date,
    nextRun: Date,
    actionIcon: JSX.Element,
    actionCallback(): void
}

export const TaskInfoCard: React.FC<TaskInfoCardProps> = ({name, prevRun, nextRun, actionIcon, actionCallback}) => {
    return (
        <Card sx={{padding: '1em'}}>
            <Typography color="text.secondary">Previous run: {getRelativeDateString(prevRun)}</Typography>
            <Typography variant="h5" color="inherit" sx={{textTransform: 'capitalize', fontWeight: "Bold"}}
                        gutterBottom>{name}</Typography>
            <Typography color="text.secondary">Next run: {getRelativeDateString(nextRun)}</Typography>
            <Grid container direction="row-reverse">
                <Grid item>

                    <Fab color="secondary" aria-label="edit" onClick={actionCallback}>
                        {actionIcon}
                    </Fab>
                </Grid>
            </Grid>
        </Card>
    );
};