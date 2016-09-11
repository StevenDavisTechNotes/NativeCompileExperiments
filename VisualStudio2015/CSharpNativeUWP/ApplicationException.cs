using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace SciMark2
{
    public class ApplicationException: Exception
    {
        public ApplicationException(string msg): base(msg) { }
    }
}
