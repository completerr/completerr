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
import Paper from "@mui/material/Paper";
import Head from 'next/head';

const theme = createTheme({
    palette: {
        primary: teal
    }
});
const NavIcon = styled(CompleterrIcon)`
  height: 2em;
  padding-right: 0.5em;
  display: inline-block;
  vertical-align: middle;
`
const NavTitle = styled.span`
  font-family: Poppins-Bold;
  text-transform: uppercase;
  margin-right: 3em;
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
            <Head>
                <title>Completerr</title>
            </Head>
            <CssBaseline/>
            <AppBar position="relative" style={{background: '#000'}}>
                <Toolbar>
                    <Box sx={{display: {xs: 'flex'}}}>
                        <Typography variant="h6" noWrap>
                            <NavIcon/>

                            <NavTitle>Completerr</NavTitle>
                        </Typography>
                    </Box>
                    <Box sx={{flexGrow: 1, display: {xs: 'flex'}}}>
                        <Button sx={{my: 2, color: 'white', display: 'block'}} onClick={() => {
                            setCurrentNav(Nav.TaskHistory)
                        }}>Task </Button>
                        <Button sx={{my: 2, color: 'white', display: 'block'}} onClick={() => {
                            setCurrentNav(Nav.RadarrHistory)
                        }}>Radarr </Button>
                        <Button sx={{my: 2, color: 'white', display: 'block'}} onClick={() => {
                            setCurrentNav(Nav.SonarrHistory)
                        }}>Sonarr </Button>
                    </Box>
                </Toolbar>
            </AppBar>

            <Box component={Paper}
                 sx={{
                     backgroundColor: (theme) =>
                         theme.palette.mode === 'light'
                             ? theme.palette.grey[100]
                             : theme.palette.grey[900],
                     flexGrow: 1,
                 }}>
                <Container sx={{mt: 4, mb: 4}}>
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
export default Home
