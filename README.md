# web-service-gin

router.GET("/getuser", m.GetUsers) <br/>
router.GET("/userid/:id", m.GetUsersByID)  <br/>
router.GET("/getAllProducts", m.GetAllProducts) <br/>
router.GET("/getProduct/:id", m.GetProductByID) <br/>
router.GET("/getRewardByUserId/:id", m.GetRewardByUserId) <br/>
router.GET("/getReceipt", m.GetReceipt) <br/>
router.GET("/getReceipt/:id", m.GetReceiptByUserID) <br/>
router.POST("/insertUser", m.InsertUser) <br/>
router.POST("/updateUser", m.UpdateUser) <br/>
router.POST("/addUserPoint", m.UpdateUserPoint) <br/>
router.POST("/insertProduct", m.InsertProduct) <br/>
router.POST("/insertReward", m.InsertReward) <br/>
router.POST("/uploadReceipt", m.Upload) <br/>
