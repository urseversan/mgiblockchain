<%@ taglib uri="/struts-tags" prefix="s" %> 
<!DOCTYPE html>
<!--[if lt IE 7]><html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]><html class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]><html class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!--> <html class="no-js"> <!--<![endif]-->
<head>
<!-- 
Kool Store Template
http://www.templatemo.com/preview/templatemo_428_kool_store
-->
    <meta charset="utf-8">
    <title>Future Clearance and Settlement Group</title>

    <meta name="description" content="">
    <meta name="viewport" content="width=device-width">

    <link href="http://fonts.googleapis.com/css?family=Open+Sans:400,300,600,700,800" rel="stylesheet">

    <link rel="stylesheet" href="css/bootstrap.css">
    <link rel="stylesheet" href="css/normalize.min.css">
    <link rel="stylesheet" href="css/font-awesome.min.css">
    <link rel="stylesheet" href="css/animate.css">
    <link rel="stylesheet" href="css/templatemo-misc.css">
    <link rel="stylesheet" href="css/templatemo-style.css">

    <script src="js/vendor/modernizr-2.6.2.min.js"></script>

</head>
<body>
    <!--[if lt IE 7]>
    <p class="chromeframe">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> or <a href="http://www.google.com/chromeframe/?redirect=true">activate Google Chrome Frame</a> to improve your experience.</p>
    <![endif]-->

    
    <header class="site-header">
        <div class="top-header">
            <div class="container">
                <div class="row">
                    <div class="col-md-6 col-sm-6">
                        <div class="top-header-left">
                             <a href="logout.action">Log Out</a>
                        </div> <!-- /.top-header-left -->
                    </div> <!-- /.col-md-6 -->
                    <div class="col-md-6 col-sm-6">
                        <div class="social-icons">
                            <ul>
                                <li><a href="#" class="fa fa-facebook"></a></li>
                                <li><a href="#" class="fa fa-dribbble"></a></li>
                                <li><a href="#" class="fa fa-twitter"></a></li>
                                <li><a href="#" class="fa fa-linkedin"></a></li>
                            </ul>
                            <div class="clearfix"></div>
                        </div> <!-- /.social-icons -->
                    </div> <!-- /.col-md-6 -->
                </div> <!-- /.row -->
            </div> <!-- /.container -->
        </div> <!-- /.top-header -->
        <div class="main-header">
            <div class="container">
                <div class="row">
                    <div class="col-md-12 col-xs-8">
                        <div class="logo">
                            <h1><a href="home.action">Future Clearance and Settlement Group</a></h1>
                        </div> <!-- /.logo -->
                    </div> <!-- /.col-md-4 -->
                   
                </div> <!-- /.row -->
            </div> <!-- /.container -->
        </div> <!-- /.main-header -->
        <div class="main-nav">
             <div class="container">
	                <div class="row">
	                    <div class="col-md-8 col-sm-4">
	                        <div class="list-menu">
	                            <ul>
	                                <li><a href="home.action">Home</a></li>
	                                <s:if test="%{#session.login ==true}">
	                                	 <s:if test="%{#session.role =='caadmin'}">
		                                	<li><a href="members.action">Group Members</a></li>
		                                	<li><a href="viewchain.action">View Chain</a></li>
		                                 </s:if>
		                              	 <s:if test="%{#session.role !='caadmin'}">
		                              		<li><a href="audit.action">Audit Ledger</a></li>
		                                 </s:if>			                                
		                            </s:if>    
	                            </ul>
	                        </div> <!-- /.list-menu -->
	                    </div> <!-- /.col-md-6 -->
	                    
	                </div> <!-- /.row -->
	            </div> <!-- /.container -->
        </div> <!-- /.main-nav -->
    </header> <!-- /.site-header -->

    <div class="content-section">
        <div class="container">
            <div class="row">
                <div class="col-md-4">
                    <div class="product-item-1">
                        <div class="product-thumb">
                            <img src="images/gallery-image-1.jpg" alt="Product Title">
                        </div> <!-- /.product-thumb -->
                        <div class="product-content">
                            <h5><a href="#">Kool Shirt</a></h5>
                            <span class="tagline">Partner Name</span>
                            <span class="price">$25.00</span>
                            <p>Doloremque quo possimus quas necessitatibus blanditiis excepturi. Commodi, sunt asperiores tenetur deleniti labore!</p>
                        </div> <!-- /.product-content -->
                    </div> <!-- /.product-item -->
                </div> <!-- /.col-md-4 -->
                <div class="col-md-4">
                    <div class="product-holder">
                        <div class="product-item-2">
                            <div class="product-thumb">
                                <img src="images/featured/1.jpg" alt="Product Title">
                            </div> <!-- /.product-thumb -->
                            <div class="product-content overlay">
                                <h5><a href="#">Kooler Shirt</a></h5>
                                <span class="tagline">Partner Name</span>
                                <span class="price">$30.00</span>
                                <p>Doloremque quo possimus quas necessitatibus blanditiis excepturi. Commodi, sunt tenetur deleniti labore!</p>
                            </div> <!-- /.product-content -->
                        </div> <!-- /.product-item-2 -->
                        <div class="product-item-2">
                            <div class="product-thumb">
                                <img src="images/featured/2.jpg" alt="Product Title">
                            </div> <!-- /.product-thumb -->
                            <div class="product-content overlay">
                                <h5><a href="#">Koolest Shirt</a></h5>
                                <span class="tagline">Partner Name</span>
                                <span class="price">$45.00</span>
                                <p>Doloremque quo possimus quas necessitatibus blanditiis excepturi. Commodi, sunt tenetur deleniti labore!</p>
                            </div> <!-- /.product-content -->
                        </div> <!-- /.product-item-2 -->
                        <div class="clearfix"></div>
                    </div> <!-- /.product-holder -->
                </div> <!-- /.col-md-4 -->
                <div class="col-md-4">
                    <div class="product-item-3">
                        <div class="product-thumb">
                            <img src="images/featured/6.jpg" alt="">
                        </div> <!-- /.product-thumb -->                        
                    </div> <!-- /.product-item-3 -->
                </div> <!-- /.col-md-4 -->
            </div> <!-- /.row -->            
        </div> <!-- /.container -->
    </div> <!-- /.content-section -->
    
    <!-- Dummy data will be removed when integrated with JSP changes -->
    <s:iterator value="membersList">
		<p>
			Members : 	 <s:property value="memberName"/> ,
				 		 <s:property value="memberLogoFileName"/>,
				 		 <s:property value="memberDetails"/> 
		</p>
		<br></br>
	</s:iterator>

 <footer class="site-footer">        
        <div class="bottom-footer">
            <div class="container">
                <div class="row">
                    <div class="col-md-12 text-center">
                        <span>Copyright &copy; 2016 <a href="home.action">FUTURE CLEARANCE AND SETTLEMENT GROUP</a></span>
                    </div> <!-- /.col-md-12 -->
                </div> <!-- /.row -->
            </div> <!-- /.container -->
        </div> <!-- /.bottom-footer -->
 </footer> <!-- /.site-footer -->

    
    <script src="js/vendor/jquery-1.10.1.min.js"></script>
    <script>window.jQuery || document.write('<script src="js/vendor/jquery-1.10.1.min.js"><\/script>')</script>
    <script src="js/jquery.easing-1.3.js"></script>
    <script src="js/bootstrap.js"></script>
    <script src="js/plugins.js"></script>
    <script src="js/main.js"></script>


</body>
</html>