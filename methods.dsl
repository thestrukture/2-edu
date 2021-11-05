<gos> 
 
<method name="IsActive" var="url string, lookup string" return="string"  >
    return "active"
</method>

<method name="GetState" var="r *http.Request" return="PageComp"  >
    return PageComp{ "/" }
</method>

 </gos> 