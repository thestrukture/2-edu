<gos> 
 
<method name="IsActive" var="url, lookup" return="string"  >
    return strings.Contains(url.(string), lookup.(string))
</method>

<method name="GetState" var="" return="PageComp"  >
    return PageComp{ State : "/" }
</method>

 </gos> 