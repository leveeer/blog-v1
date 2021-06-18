/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/light");

var $root = ($protobuf.roots["default"] || ($protobuf.roots["default"] = new $protobuf.Root()))
.setOptions({
  go_package: "../go_proto;proto"
})
.addJSON({
  proto: {
    nested: {
      CsId: {
        values: {
          CsBeginIndex: 0,
          CsGetArticles: 1,
          CsGetArticleById: 2,
          CsGetBlogHomeInfo: 3
        }
      },
      RequestPkg: {
        fields: {
          cmdId: {
            type: "CsId",
            id: 1
          },
          currentPage: {
            type: "uint32",
            id: 2
          },
          articleId: {
            type: "uint32",
            id: 3
          }
        }
      },
      Article: {
        fields: {
          id: {
            type: "int32",
            id: 1
          },
          userId: {
            type: "int32",
            id: 2
          },
          categoryID: {
            type: "int32",
            id: 3
          },
          articleCover: {
            type: "string",
            id: 4
          },
          articleTitle: {
            type: "string",
            id: 5
          },
          articleContent: {
            type: "string",
            id: 6
          },
          createTime: {
            type: "int64",
            id: 7
          },
          updateTime: {
            type: "int64",
            id: 8
          },
          isTop: {
            type: "bool",
            id: 9
          },
          isPublish: {
            type: "bool",
            id: 10
          },
          isDelete: {
            type: "bool",
            id: 11
          },
          isOriginal: {
            type: "bool",
            id: 12
          },
          clickCount: {
            type: "int64",
            id: 13
          },
          collectCount: {
            type: "int64",
            id: 14
          },
          tags: {
            rule: "repeated",
            type: "Tag",
            id: 15
          },
          categoryName: {
            type: "string",
            id: 16
          }
        }
      },
      Tag: {
        fields: {
          id: {
            type: "int32",
            id: 1
          },
          tagName: {
            type: "string",
            id: 2
          },
          createTime: {
            type: "int64",
            id: 3
          },
          updateTime: {
            type: "int64",
            id: 4
          },
          status: {
            type: "bool",
            id: 5
          },
          clickCount: {
            type: "int64",
            id: 6
          }
        }
      },
      BlogHomeInfo: {
        fields: {
          userInfo: {
            type: "UserInfo",
            id: 1
          },
          articleCount: {
            type: "int64",
            id: 2
          },
          categoryCount: {
            type: "int64",
            id: 3
          },
          tagCount: {
            type: "int64",
            id: 4
          },
          notice: {
            type: "string",
            id: 5
          },
          viewCount: {
            type: "int64",
            id: 6
          }
        }
      },
      UserInfo: {
        fields: {
          id: {
            type: "int32",
            id: 1
          },
          email: {
            type: "string",
            id: 2
          },
          nickName: {
            type: "string",
            id: 3
          },
          avatar: {
            type: "string",
            id: 4
          },
          intro: {
            type: "string",
            id: 5
          },
          website: {
            type: "string",
            id: 6
          },
          createTime: {
            type: "int64",
            id: 7
          },
          updateTime: {
            type: "int64",
            id: 8
          },
          isDisable: {
            type: "bool",
            id: 9
          }
        }
      },
      ArticleInfo: {
        fields: {
          article: {
            type: "Article",
            id: 1
          },
          lastArticle: {
            type: "Article",
            id: 2
          },
          nextArticle: {
            type: "Article",
            id: 3
          },
          recommendArticleList: {
            rule: "repeated",
            type: "Article",
            id: 4
          },
          articleLatestList: {
            rule: "repeated",
            type: "Article",
            id: 5
          }
        }
      },
      Archives: {
        fields: {
          archiveList: {
            rule: "repeated",
            type: "ArchiveArticleInfo",
            id: 1
          },
          count: {
            type: "int32",
            id: 2
          }
        }
      },
      ArchiveArticleInfo: {
        fields: {
          id: {
            type: "int32",
            id: 1
          },
          articleTitle: {
            type: "string",
            id: 2
          },
          createTime: {
            type: "int64",
            id: 3
          }
        }
      },
      Category: {
        fields: {
          id: {
            type: "int32",
            id: 1
          },
          categoryName: {
            type: "string",
            id: 2
          },
          articleCount: {
            type: "int32",
            id: 3
          }
        }
      },
      ArticlesByCategoryOrTag: {
        fields: {
          articleList: {
            rule: "repeated",
            type: "Article",
            id: 1
          },
          name: {
            type: "string",
            id: 2
          }
        }
      },
      Response: {
        values: {
          ResponseBeginIndex: 0
        }
      },
      ResultCode: {
        values: {
          Success: 0,
          Fail: 1,
          SuccessOK: 10000
        }
      },
      ResponsePkg: {
        fields: {
          cmdId: {
            type: "Response",
            id: 1
          },
          code: {
            type: "ResultCode",
            id: 2
          },
          categories: {
            rule: "repeated",
            type: "Category",
            id: 3
          },
          articlesByCategoryOrTag: {
            type: "ArticlesByCategoryOrTag",
            id: 4
          },
          tags: {
            rule: "repeated",
            type: "Tag",
            id: 5
          },
          errMsg: {
            type: "string",
            id: 10
          },
          serverTime: {
            type: "int64",
            id: 11
          },
          articleList: {
            rule: "repeated",
            type: "Article",
            id: 12
          },
          blogHomeInfo: {
            type: "BlogHomeInfo",
            id: 13
          },
          articleInfo: {
            type: "ArticleInfo",
            id: 14
          },
          archiveInfo: {
            type: "Archives",
            id: 15
          }
        }
      }
    }
  }
});

module.exports = $root;
