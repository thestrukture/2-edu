<gos> 
 
<method name="IsActive" var="url, lookup" return="string"  >
    if strings.Contains(url.(string), lookup.(string)) {
        
        return "active"
    }
    
    return ""
</method>

<method name="GetState" var="request" return="PageComp"  >

    r := request.(*http.Request)
    
    return PageComp{ State : r.RequestURI }
</method>

 </gos> 