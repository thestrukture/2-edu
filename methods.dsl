<gos> 
 
<method name="IsActive" var="url, lookup" return="string"  >
    if strings.Contains(url.(string), lookup.(string)) {
        
        return "active"
    }
    
    return ""
</method>

<method name="GetState" var="" return="PageComp"  >
    return PageComp{ State : "/" }
</method>

 </gos> 