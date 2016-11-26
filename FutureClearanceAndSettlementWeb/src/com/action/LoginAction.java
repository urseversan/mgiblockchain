/**
 * 
 */
package com.action;

/**
 * @author Niranjan
 *
 */
import com.opensymphony.xwork2.Action;

import java.util.Map;

import org.apache.struts2.dispatcher.SessionMap;  
import org.apache.struts2.interceptor.SessionAware;

public class LoginAction implements Action, SessionAware {    

    //Java Bean to hold the form parameters
    private String username;
    private String password;
    SessionMap<String,String> sessionmap;  

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
    
    @Override
    public String execute() throws Exception {
        if (validateString(getUsername()) && validateString(getPassword()))
            return "SUCCESS";
        else return "ERROR";
    }
    
	@SuppressWarnings("unchecked")
	@Override
	public void setSession(@SuppressWarnings("rawtypes") Map map) {
		sessionmap=(SessionMap)map;  
	    sessionmap.put("login","true"); 		
	}
	
	public String logout(){  
	    sessionmap.invalidate();  
	    return "success";  
	}  
}
