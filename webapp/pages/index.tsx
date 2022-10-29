import type {NextPage} from 'next'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import {SearchHistory} from "../components/search/history";
import AppBar from '@mui/material/AppBar';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import {createTheme, ThemeProvider,} from '@mui/material/styles';
import {teal} from '@mui/material/colors';
import {CompleterrIcon} from "../components/Icon";
import styled from "styled-components";
import {useState} from "react";
import {TaskHistory} from "../components/task/history";
import Button from '@mui/material/Button';

const theme = createTheme({
    palette: {
        primary: teal
    }
});
const NavIcon = styled(CompleterrIcon)`
  height: 1.5em;
  padding-right: 0.5em;
  display: inline-block;
  vertical-align: middle;
`

enum Nav {
    TaskHistory = 0,
    RadarrHistory,
    SonarrHistory,
}

const getPage = (selection: Nav): JSX.Element => {
    switch (selection) {
        case Nav.TaskHistory:
            return <TaskHistory/>
        case Nav.RadarrHistory:
            return <SearchHistory type={"radarr"}/>
        case Nav.SonarrHistory:
            return <SearchHistory type={"sonarr"}/>
    }
}
const Home: NextPage = () => {
    const [currentNav, setCurrentNav] = useState(Nav.TaskHistory)
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline/>
            <AppBar position="relative" color="primary">
                <Toolbar>
                    <Typography variant="h6" noWrap>
                        <NavIcon/>
                        Completerr

                    </Typography>
                    <Box sx={{flexGrow: 1, display: {xs: 'flex'}}}>
                        <Button sx={{ my: 2, color: 'white', display: 'block' }} onClick={() => {
                            setCurrentNav(Nav.TaskHistory)
                        }}>Task </Button>
                        <Button sx={{ my: 2, color: 'white', display: 'block' }} onClick={() => {
                            setCurrentNav(Nav.RadarrHistory)
                        }}>Radarr </Button>
                        <Button sx={{ my: 2, color: 'white', display: 'block'}}  onClick={() => {
                            setCurrentNav(Nav.SonarrHistory)
                        }}>Sonarr </Button>
                    </Box>
                </Toolbar>
            </AppBar>
            <Box
                sx={{
                    backgroundColor: (theme) =>
                        theme.palette.mode === 'light'
                            ? theme.palette.grey[100]
                            : theme.palette.grey[900],
                    flexGrow: 1,
                }}>
                <Container>
                    {
                        getPage(currentNav)
                    }
                </Container>
                {/* Footer */}
                <Box sx={{bgcolor: 'background.paper', p: 6}} component="footer">
                    <Typography variant="h6" align="center" gutterBottom>
                        <a
                            href="https://github.com/completerr/completerr"
                            target="_blank"
                            rel="noopener noreferrer"
                        >
                            Contribute on{' '}
                            <span className={styles.logo}>
                       <Image src="/github.svg" alt="Github Logo" width={16} height={16}/>
                               </span></a>
                    </Typography>
                </Box>
            </Box>
            {/* End footer */}
        </ThemeProvider>
    );
}
// const Home: NextPage = () => {
//     return (
//         <div className={styles.container}>
//             <Head>
//                 <title>Completerr</title>
//                 <link rel="icon" href="/favicon.ico"/>
//             </Head>
//
//             <Navbar/>
//             <main>
//                 <div className={styles.containerTable100}>
//                     <div className={styles.wrapTable100}>
//                         {/*<TaskHistory/>*/}
//                         <SearchHistory type={"radarr"}/>
//                         <SearchHistory type={"sonarr"}/>
//                     </div>
//                 </div>
//             </main>
//
//             <footer className={styles.footer}>
//                 <a
//                     href="https://github.com/completerr/completerr"
//                     target="_blank"
//                     rel="noopener noreferrer"
//                 >
//                     Contribute on{' '}
//                     <span className={styles.logo}>
//             <Image src="/github.svg" alt="Github Logo" width={16} height={16}/>
//           </span>
//                 </a>
//             </footer>
//         </div>
//     )
// }
//
export default Home
