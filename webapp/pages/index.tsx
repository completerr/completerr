import type {NextPage} from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import {Navbar} from "../components/nav/navbar";
import {TaskHistory} from "../components/task/history";
import {SearchHistory} from "../components/search/history";

const Home: NextPage = () => {
    return (
        <div className={styles.container}>
            <Head>
                <title>Completerr</title>
                <link rel="icon" href="/favicon.ico"/>
            </Head>

            <Navbar/>
            <main>
                <div className={styles.containerTable100}>
                    <div className={styles.wrapTable100}>
                        {/*<TaskHistory/>*/}
                        <SearchHistory type={"radarr"}/>
                        <SearchHistory type={"sonarr"}/>
                    </div>
                </div>
            </main>

            <footer className={styles.footer}>
                <a
                    href="https://github.com/completerr/completerr"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Contribute on{' '}
                    <span className={styles.logo}>
            <Image src="/github.svg" alt="Github Logo" width={16} height={16}/>
          </span>
                </a>
            </footer>
        </div>
    )
}

export default Home
