using System;
namespace LeetcodeStudy.Solutions.LocationAttrbutes
{
    public class LocationAttribute:Attribute
    {
        public String Url;
        public LocationAttribute(string s)
        {
            Url = s;
        }
    }

    public class DescriptionAttribute : Attribute
    {
        public string Description;

        public DescriptionAttribute(string d)
        {
            Description = d;
        }
    }
}