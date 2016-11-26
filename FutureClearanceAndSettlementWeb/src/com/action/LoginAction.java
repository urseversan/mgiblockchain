/**
 * 
 */
package com.action;

/**
 * @author Niranjan
 *
 */
import com.opensymphony.xwork2.Action;
import com.opensymphony.xwork2.ActionSupport;
import com.opensymphony.xwork2.ActionInvocation;

import java.util.Map;

import org.apache.struts2.dispatcher.SessionMap;  
import org.apache.struts2.interceptor.SessionAware;

public class LoginAction extends ActionSupport implements Action, SessionAware{    

    //Java Bean to hold the form parameters
    private String username;
    private String password;
    SessionMap<String,String> sessionmap;
    private boolean login;

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    private boolean validateString(String str) {
        if (str != null && !str.equals(""))
            return true;
        return false;
    }    
    
    public boolean isLogin() {
		return login;
	}

	public void setLogin(boolean login) {
		this.login = login;
	}

	private boolean validateLogin(String username, String password){
    	if (validateString(username) && validateString(password)){
    		if(username.equalsIgnoreCase("blockchain") && password.equalsIgnoreCase("blockchain")){
    			return true;
    		}  
    	}
    	return false;
    }
    
    @Override
    public String execute() throws Exception {
        if (validateLogin(getUsername(),getPassword())){
        	addActionMessage("Welcome user!");
        	setLogin(true);        	
        	return "SUCCESS";
        }            
        else{
        	addActionError("Invalid Username or password");
        	return "ERROR";
        } 
    }
    
	@SuppressWarnings("unchecked")
	public void setSession(Map map) {		
		sessionmap=(SessionMap)map;  
	    sessionmap.put("login","true"); 		
	}
	
	public String logout(){
		setLogin(false);
	    sessionmap.invalidate();  
	    return "SUCCESS";  
	}  
}
