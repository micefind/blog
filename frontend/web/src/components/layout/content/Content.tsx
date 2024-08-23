import styles from "./content.module.scss"
import logo from "../../../../public/logo.svg"
import { List, Tag } from "antd"
import { FolderViewOutlined } from "@ant-design/icons"
export default function Content() {
  const list = []
  for (let i = 0; i < 20; i++) {
    list.push({ title: i })
  }
  return (
    <div className={styles.content}>
      {/* 头部 */}
      <div className={styles.header}>
        <div className={styles.logo}>
          <a href="/">
            <img src={logo} alt="鼠觅奇物" />
            <h1>鼠觅奇物</h1>
          </a>
        </div>
      </div>
      {/* 内容 */}
      <div className={styles.content}>
        <div className={styles.list}>
          <List
            itemLayout="horizontal"
            dataSource={list}
            renderItem={() => (
              <List.Item>
                <div className={styles.item}>
                  <div className={styles.text}>
                    <h2>
                      可视化大屏开发，知道这些解决方案，效率至少提升2倍！（中）
                    </h2>
                    <div className={styles.describe}>
                      <p>
                        Hello大家好，我是日拱一卒的攻城师不浪，专注前端、后端、AI学习、二三维可视化、GIS等学习沉淀，这是2024年输出的第15/100篇文章，欢迎志同道合的朋友交流合作；
                      </p>
                    </div>
                    <div className={styles.tag}>
                      <div>
                        <span>2024/08/16</span>
                        <span>
                          <FolderViewOutlined />
                          998
                        </span>
                      </div>
                      <div>
                        <Tag color="default">标签1</Tag>
                        <Tag color="default">标签版本1</Tag>
                        <Tag color="default">标签1</Tag>
                      </div>
                    </div>
                  </div>
                  <div className={styles.img}></div>
                </div>
              </List.Item>
            )}
          />
        </div>
      </div>
    </div>
  )
}
