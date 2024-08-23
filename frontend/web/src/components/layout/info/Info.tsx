import styles from "./info.module.scss"
import github from "../../../assets/images/github.png"
import gitee from "../../../assets/images/gitee.png"
import bilibili from "../../../assets/images/bilibili.png"

export default function Info() {
  const year = new Date().getFullYear()
  return (
    <div className={styles.info}>
      <div className={styles.profile}>
        <div className={styles.avatar}>
          <img
            src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
            alt=""
          />
        </div>
        <div className={styles.name}>
          <h2>micefind</h2>
        </div>
        <div className={styles.icon}>
          <a href="https://github.com/micefind" target="_blank">
            <img src={github} alt="github" />
          </a>
          <a
            href="https://gitee.com/micefind"
            target="_blank"
            className={styles.gitee}
          >
            <img src={gitee} alt="gitee" />
          </a>
          <a
            href="https://space.bilibili.com/321496512?spm_id_from=333.788.0.0"
            target="_blank"
          >
            <img src={bilibili} alt="bilibili" />
          </a>
        </div>
      </div>
      <div className={styles.list}>
        <div className={styles.content}>
          <div></div>
          <ul>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
            <a href="https://github.com/micefind/museum" target="_blank">
              <li>
                <div className={styles.img}>
                  <img
                    src="https://avatars.githubusercontent.com/u/109268208?v=4&size=64"
                    alt="线上博物馆"
                  />
                </div>
                <div className={styles.text}>
                  <h3>线上博物馆</h3>
                  <span>dsf</span>
                </div>
              </li>
            </a>
          </ul>
        </div>
        <div className={styles.footer}>
          <div>
            <span>© {year} MiceFind.com</span>
          </div>
          <div>
            <a href="https://beian.miit.gov.cn/" target="_blank">
              <span>渝ICP备2022007989号</span>
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}
