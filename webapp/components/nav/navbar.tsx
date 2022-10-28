import styles from '../../styles/Home.module.css'

interface NavButtonProps  {
    label: string
    onClick() :void
}
type NavbarProps = {
    buttons?: NavButtonProps[]
}
export const Navbar: React.FC<NavbarProps> = ({ buttons}:NavbarProps )=> {
    return (
        <div className={styles.navContainer}>
            <>
                <div>Completerr</div>
                { buttons && buttons.map(b =>{
                    <button onClick={b.onClick} >{b.label}</button>
                })}
            </>
        </div>
    );
};