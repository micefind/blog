import Content from "../../components/layout/content/Content";
import Info from "../../components/layout/info/Info";
import styles from "./layout.module.scss";

export default function Home() {
  return (
    <div className={styles.layout}>
      <div className={styles.content}>
        <Content />
      </div>
      <div className={styles.info}>
        <Info />
      </div>
    </div>
  );
}
