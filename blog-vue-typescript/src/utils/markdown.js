import {convert} from "utils/utils";

const highlight = require("highlight.js");
const marked = require("marked");
const tocObj = {
    toc: [],
    index: 0,
    add(text, level) {
        const anchor = `#toc${level}${++this.index}`;
        this.toc.push({anchor: anchor, level: level, text: text});
        return anchor;
    },

    getToc() {
        //return this.toc
        let tocArr = [];
        let idx = 0
        for (let i = 0; i < this.toc.length; i++) {
            tocArr.push({
                id: i + 1,
                parentId: 0,
                anchor: this.toc[i].anchor,
                level: this.toc[i].level,
                text: this.toc[i].text,
                subTitle: [],
                active: false,
            })
        }
        for (let i = 1; i < tocArr.length; i++) {
            if (tocArr[idx].level < tocArr[i].level) {
                tocArr[i].parentId = tocArr[idx].id;
            }
            if (tocArr[idx].level === tocArr[i].level) {
                tocArr[i].parentId = tocArr[idx].parentId;
            }

            if (tocArr[idx].level > tocArr[i].level) {
                for (let j = i; j > 0; j--) {
                    if (tocArr[j].level < tocArr[i].level) {
                        tocArr[i].parentId = tocArr[j].id;
                        break;
                    }
                }
            }
            idx++;
        }
        return convert(tocArr)
    },
};

class MarkUtils {
    constructor() {
        this.rendererMD = new marked.Renderer();
        this.rendererMD.heading = function (text, level, raw) {
            const anchor = tocObj.add(text, level);
            return `<h${level} id=${anchor}>${text}</h${level}>\n`;
        };
        this.rendererMD.table = function (header, body) {
            return '<table class="table">' + header + body + '</table>'
        }

        highlight.configure({useBR: true});
        marked.setOptions({
            renderer: this.rendererMD,
            headerIds: false,
            gfm: true,
            tables: true,
            breaks: false,
            pedantic: false,
            sanitize: false,
            smartLists: true,
            smartypants: false,
            highlight: function (code) {
                return highlight.highlightAuto(code).value;
            }
        });
    }

    async marked(data) {
        if (data) {
            return {content: await marked(data), toc: tocObj.getToc()};
        } else {
            return null;
        }
    }
}

const markdown = new MarkUtils();

export default markdown;
